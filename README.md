go-genums
=========

Utility to generate type-checked enums, friendly for ``go:generate`` syntax.

Usage
=====

```
go-genums Wolf piggy int examples/intEnums/intEnums.go
```

* ``Wolf`` is the enum prefix used in generated output code
* ``examples/intEnums/intEnums.go`` is the source code containing the Go enum declaration (you need to write this)
* ``int`` is the type of the Go enum declaration in above user-provided source
* ``piggy`` is the internal enum prefix in the user-provided source

The generated code will contain:

* an interface type for the enum e.g. ``WolfEnum`` in this example
* a func returning a slice of all valid enum types ``WolfEnumValues()`` in this example
* a factory method to cast values into a valid enum, ``NewWolfFromValue(v int)`` in this example
* a set of struct types (one for each enum value) that all satisfy the enum interface (``WolfEnum``) and implement methods ``String()`` (enum name) and ``Value()``

Furthermore, comparing two enum structs will work as expected; see [examples/intEnums/main.go](examples/intEnums/main.go) for an usage example.

Examples
========

See ``examples/`` for an integer enum and string enum example.
