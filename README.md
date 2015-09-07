go-genums
=========

Utility to generate type-checked enums, friendly for usage with ``go:generate``.

Usage
=====

Example:
```
genums Day day int examples/dayEnum/main.go
```

In the above command line:
* ``Day`` is the prefix of the enum interface name that will be generated e.g. ``DayEnum``
* ``day`` is the (internal, in this case) prefix used in the constant identifiers declared in the user-provided Go source
* ``int`` is the type of the declared constants in the user-provided Go source
* [examples/dayEnum/main.go](examples/dayEnum/main.go) is the user-provided Go source file name that already contains a Go const enum

The generated code will provide:

* an interface type for the enum, ``DayEnum``
* a func returning a slice of all valid enum values, ``DayEnumValues()``
* a factory method to cast values into a valid enum, ``NewDayFromValue(v int)``
* a factory method that exclusively casts values into a valid enum and panics otherwise, ``MustGetDayFromValue(v int)``
* a set of struct types (one for each legit enum value) that all satisfy the enum interface, ``DayEnum``

## Features overview

Comparing two enum values will yield the intuitively expected result:
```
a := Monday{}.New()
b := Monday{}.New()

fmt.Println(a == b) // prints true
```

You can use a type switch to evaluate the interface-typed enum values:
```
switch a.(type) {
	case Monday:
		fmt.Println("Found the right day")
	default:
		panic("Cannot find the right day!")
}
```

If you put the generated code in a package the ``.value`` field becomes inaccessible, thus all enum values will be immutable and even safer to use.

Each enum value has a ``String()`` method that returns its descriptive name (same as Go identifier) and a ``Value()`` method that returns the correspondent value defined in the user-provided source.

How to not use
==============

Do not use the enum types without calling their New() method:
```
a := Tuesday{}
b := Friday{}

fmt.Println(a == b) // returns true

fmt.Println(Tuesday.Value()) // returns 0, which is the default for any uninitialised variable in Go
```

Preferably, put your enums in a package so that they are sealed and you will not be able to inadvertenty do something like:
```
a.value++
```

Examples
========

See ``examples/`` for integer, string and struct examples.

License
=======

go-genums is licensed under an [MIT license](LICENSE.md).
