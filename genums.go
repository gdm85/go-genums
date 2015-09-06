package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type genumsContext struct {
	packageName    string
	prefix         string
	internalPrefix string
	valueType      string
}

const outputTemplate = `package %s

// %sEnum is the the enum interface that can be used
type %sEnum interface {
	String() string
	Value() %s
}

// %sEnumBase is the internal, non-exported type
type %sEnumBase struct{ value %s }

// Value() returns the enum value
func (eb %sEnumBase) Value() %s { return eb.value }

// String() returns the enum name as you use it in Go code,
// needs to be overriden by inheriting types
func (eb %sEnumBase) String() string { return "" }

`

func (ge genumsContext) generateCode(valueSuffixes []string) {
	// for use in generated but non-exported types
	internalGeneratedPrefix := strings.ToLower(ge.prefix[0:1]) + ge.prefix[1:]

	// generate enum base type
	fmt.Printf(outputTemplate, ge.packageName,
		ge.prefix,
		ge.prefix,
		ge.valueType,
		internalGeneratedPrefix,
		internalGeneratedPrefix, ge.valueType,
		internalGeneratedPrefix, ge.valueType,
		internalGeneratedPrefix)

	// generate value types
	for _, suffix := range valueSuffixes {
		fmt.Printf("// %s%s is the enum type for '%s%s' value\ntype %s%s struct{ %sEnumBase }\n\n",
			ge.prefix, suffix, ge.internalPrefix, suffix,
			ge.prefix, suffix, internalGeneratedPrefix)

		fmt.Printf("// New is the constructor for a brand new %sEnum with value '%s%s'\nfunc (eb %s%s) New() %sEnum { return %s%s{%sEnumBase{%s%s}} }\n\n",
			ge.prefix, ge.internalPrefix, suffix,
			ge.prefix, suffix, ge.prefix,
			ge.prefix, suffix, internalGeneratedPrefix, ge.internalPrefix, suffix)

		fmt.Printf("// String returns always \"%s%s\" for this enum type\nfunc (eb %s%s) String() string { return \"%s%s\" }\n\n",
			ge.prefix, suffix,
			ge.prefix, suffix,
			ge.prefix, suffix)
	}

	// generate collection of all valid values
	fmt.Printf("var internal%sEnumValues = []%sEnum{\n", ge.prefix, ge.prefix)
	for _, suffix := range valueSuffixes {
		fmt.Printf("\t%s%s{}.New(),\n", ge.prefix, suffix)
	}
	fmt.Printf("}\n\n// %sEnumValues will return all available enum value types\nfunc %sEnumValues() []%sEnum { return internal%sEnumValues[:] }\n\n",
		ge.prefix, ge.prefix, ge.prefix, ge.prefix)

	// generate the factory func that initializes by value
	fmt.Printf(`// New%sFromValue will generate a valid enum from a value, or return nil in case of invalid value
func New%sFromValue(v %s) (result %sEnum) {
	switch v {
`, ge.prefix, ge.prefix, ge.valueType, ge.prefix)

	for _, suffix := range valueSuffixes {
		fmt.Printf("\tcase %s%s:\n\t\tresult = %s%s{}.New()\n", ge.internalPrefix, suffix, ge.prefix, suffix)
	}
	fmt.Printf("\t}\n\treturn\n}\n")

	// add a factory method that panics
	fmt.Printf(`
// MustGet%sFromValue is the same as New%sFromValue, but will panic in case of conversion failure
func MustGet%sFromValue(v %s) %sEnum {
	result := New%sFromValue(v)
	if result == nil {
		panic("invalid %sEnum value cast")
	}
	return result
}
`, ge.prefix, ge.prefix,
		ge.prefix, ge.valueType, ge.prefix,
		ge.prefix, ge.prefix)
}

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "Usage: genums EnumPrefix internalEnumPrefix internalEnumValueType sourceEnum.go\n")
		fmt.Fprintf(os.Stderr, "You can specify a valid Go enum in enums.go and genums will parse it based on the internal prefix you specified\n")
		os.Exit(1)
		return
	}

	generatedPrefix := os.Args[1]
	if len(generatedPrefix) < 2 {
		fmt.Fprintf(os.Stderr, "genums: invalid generated prefix specified, should be at least 2 characters long\n")
		os.Exit(1)
		return
	}

	internalPrefix := os.Args[2]
	if len(internalPrefix) == 0 {
		fmt.Fprintf(os.Stderr, "genums: invalid internal prefix specified\n")
		os.Exit(1)
		return
	}

	internalType := os.Args[3]
	if len(internalType) == 0 {
		fmt.Fprintf(os.Stderr, "genums: invalid internal enum value type specified\n")
		os.Exit(1)
		return
	}

	source, err := ioutil.ReadFile(os.Args[4])
	if err != nil {
		fmt.Fprintf(os.Stderr, "genums: could not read source enums: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// individuate package name
	packageRx := regexp.MustCompile("(?m)^\\s*package\\s+([^\\s]+)$")
	pm := packageRx.FindSubmatch(source)
	if len(pm) == 0 {
		fmt.Fprintf(os.Stderr, "genums: could not individuate package name\n")
		os.Exit(1)
		return
	}
	pkgName := string(pm[1])

	// match all enum values declarations
	evRx := regexp.MustCompile("(?m)^\\s*" + internalPrefix + "([^\\s=]+)\\s*(=.+)?$")
	matches := evRx.FindAllSubmatch(source, -1)
	if len(matches) == 0 {
		fmt.Fprintf(os.Stderr, "genums: could not match any enum with internal prefix '%s'\n", internalPrefix)
		os.Exit(1)
		return
	}

	// collect all values definition
	// NOTE: will not work with fairly complex Go expressions e.g. multi-line
	valueSuffixes := []string{}
	for _, m := range matches {
		valueSuffixes = append(valueSuffixes, string(m[1]))
	}

	context := genumsContext{
		packageName:    pkgName,
		prefix:         generatedPrefix,
		internalPrefix: internalPrefix,
		valueType:      internalType,
	}

	context.generateCode(valueSuffixes)
}
