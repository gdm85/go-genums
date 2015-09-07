go-genums
=========

Utility to generate type-checked enums, friendly for ``go:generate`` syntax.

Usage
=====

Example:
```
go-genums Wolf piggy int examples/intEnums/intEnums.go
```

In the above command-line:
* ``Wolf`` is the prefix of the enum interface name that will be generated
* ``examples/intEnums/intEnums.go`` is the user-provided Go source code that contains a Go const enum
* ``int`` is the type of the Go enum declaration in above user-provided source
* ``piggy`` is the internal enum prefix in the user-provided source

The generated code will provide:

* an interface type for the enum, ``WolfEnum``
* a func returning a slice of all valid enum values, ``WolfEnumValues()``
* a factory method to cast values into a valid enum, ``NewWolfFromValue(v int)``
* a factory method that exclusively casts values into a valid enum and panics otherwise, ``MustGetWolfFromValue(v int)``
* a set of struct types (one for each legit enum value) that all satisfy the enum interface, ``WolfEnum`` (methods ``String()`` and ``Value``())

## Features overview

Comparing two enum values will yield the intuitively expected result:
```
a := Wolf1{}.New()
b := Wolf1{}.New()

fmt.Println(a == b) // prints true
```

You can use a type switch to evaluate the interface-typed enum values:
```
switch a.(type) {
	case Wolf1:
		fmt.Println("Found the right wolf")
	default:
		panic("Cannot find the wolf!")
}
```

If you put the generated code in a package the ``.value`` field becomes inaccessible, thus all enum values will be immutable and even safer to use.

Each enum value has a ``String()`` method that returns its descriptive name (same as Go identifier) and a ``Value()`` method that returns the correspondent value defined in the user-provided source.

See also: [examples/intEnums/main.go](examples/intEnums/main.go)

Examples
========

See ``examples/`` for an integer enum and string enum example.

License
=======

go-genums is licensed under an [MIT license](LICENSE.md).
