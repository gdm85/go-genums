package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type genumsContext struct {
	packageName    string
	prefix         string
	internalPrefix string
	valueType      string
}

const outputTemplate = `package %s

// *** generated with go-genums ***

// %sEnum is the the enum interface that can be used
type %sEnum interface {
	String() string
	Value() %s
	unique%sMethod()
}

// %sEnumBase is the internal, non-exported type
type %sEnumBase struct{ value %s }

// Value() returns the enum value
func (eb %sEnumBase) Value() %s { return eb.value }

// String() returns the enum name as you use it in Go code,
// needs to be overriden by inheriting types
func (eb %sEnumBase) String() string { return "" }

`

// generateSuffix will generate a valid suffix for an enum type name
func generateSuffix(prefix, suffix string, tinySuffix bool) string {
	if tinySuffix {
		return prefix + suffix
	}
	return strings.ToUpper(suffix[0:1]) + suffix[1:]
}

// generateCode will print the generated type enum to standard output
func (ge genumsContext) generateCode(valueSuffixes []string, tinySuffix bool) {
	// for use in generated but non-exported types
	internalGeneratedPrefix := strings.ToLower(ge.prefix[0:1]) + ge.prefix[1:]

	// generate enum base type
	fmt.Printf(outputTemplate, ge.packageName,
		ge.prefix,
		ge.prefix,
		ge.valueType,
		ge.prefix,
		internalGeneratedPrefix,
		internalGeneratedPrefix, ge.valueType,
		internalGeneratedPrefix, ge.valueType,
		internalGeneratedPrefix)

	// generate a type for each of the allowed enum values
	for _, suffix := range valueSuffixes {
		typeName := generateSuffix(ge.prefix, suffix, tinySuffix)

		fmt.Printf("// %s is the enum type for '%s%s' value\ntype %s struct{ %sEnumBase }\n\n",
			typeName, ge.internalPrefix, suffix,
			typeName, internalGeneratedPrefix)

		fmt.Printf("// New is the constructor for a brand new %sEnum with value '%s%s'\nfunc (%s) New() %sEnum { return %s{%sEnumBase{%s%s}} }\n\n",
			ge.prefix, ge.internalPrefix, suffix,
			typeName, ge.prefix,
			typeName, internalGeneratedPrefix, ge.internalPrefix, suffix)

		fmt.Printf("// String returns always \"%s\" for this enum type\nfunc (%s) String() string { return \"%s\" }\n\n",
			typeName,
			typeName,
			typeName)

		fmt.Printf("// unique%sMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature\nfunc (%s) unique%sMethod() {}\n\n",
			ge.prefix, typeName, ge.prefix)
	}

	// generate array with each of the allowed enum values, plus a function
	// to get all of them at once
	fmt.Printf("var internal%sEnumValues = []%sEnum{\n", ge.prefix, ge.prefix)
	for _, suffix := range valueSuffixes {
		typeName := generateSuffix(ge.prefix, suffix, tinySuffix)
		fmt.Printf("\t%s{}.New(),\n", typeName)
	}
	fmt.Printf("}\n\n// %sEnumValues will return a slice of all allowed enum value types\nfunc %sEnumValues() []%sEnum { return internal%sEnumValues[:] }\n\n",
		ge.prefix, ge.prefix, ge.prefix, ge.prefix)

	// generate the factory function that initializes a valid enum by value
	fmt.Printf(`// New%sFromValue will generate a valid enum from a value, or return nil in case of invalid value
func New%sFromValue(v %s) (result %sEnum) {
	switch v {
`, ge.prefix, ge.prefix, ge.valueType, ge.prefix)

	for _, suffix := range valueSuffixes {
		typeName := generateSuffix(ge.prefix, suffix, tinySuffix)
		fmt.Printf("\tcase %s%s:\n\t\tresult = %s{}.New()\n", ge.internalPrefix, suffix, typeName)
	}
	fmt.Printf("\t}\n\treturn\n}\n")

	// add a factory method that panics in case of invalid initialisation
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
		fmt.Fprintf(os.Stderr, "Usage: genums EnumPrefix valuesPrefix valueType source.go\n")
		fmt.Fprintf(os.Stderr, "'valuesPrefix' is the prefix of the enum constants you defined in 'source.go'\n")
		os.Exit(1)
		return
	}

	enumPrefix := os.Args[1]
	if len(enumPrefix) < 2 {
		fmt.Fprintf(os.Stderr, "genums: invalid enum prefix specified, should be at least 2 characters long\n")
		os.Exit(1)
		return
	}

	valuesPrefix := os.Args[2]
	if len(valuesPrefix) == 0 {
		fmt.Fprintf(os.Stderr, "genums: invalid values prefix specified\n")
		os.Exit(1)
		return
	}

	valueType := os.Args[3]
	if len(valueType) == 0 {
		fmt.Fprintf(os.Stderr, "genums: invalid value type specified\n")
		os.Exit(1)
		return
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, os.Args[4], nil, parser.DeclarationErrors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "genums: could not load Go source: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// scan all objects declared in source file's scope, looking for constants or variables
	// that begin with 'internalPrefix' string, then save the non-prefix part (suffix)
	// note that your variables should be already initialised
	valueSuffixes := []string{}
	tinySuffix := true
	for name, obj := range f.Scope.Objects {
		if (obj.Kind == ast.Var || obj.Kind == ast.Con) && strings.HasPrefix(name, valuesPrefix) {
			suffix := obj.Name[len(valuesPrefix):]

			if len(suffix) > 1 {
				tinySuffix = false
			}

			valueSuffixes = append(valueSuffixes, suffix)
		}
	}

	// refuse to continue when no const declaration was found
	if len(valueSuffixes) == 0 {
		fmt.Fprintf(os.Stderr, "genums: no constants with prefix '%s' found\n", valuesPrefix)
		os.Exit(1)
		return
	}

	// initialise all data needed for code generation
	context := genumsContext{
		packageName:    f.Name.String(),
		prefix:         enumPrefix,
		internalPrefix: valuesPrefix,
		valueType:      valueType,
	}

	// generate code and print to stdout
	context.generateCode(valueSuffixes, tinySuffix)
}
