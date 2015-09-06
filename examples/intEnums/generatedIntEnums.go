package main

// WolfEnum is the the enum interface that can be used
type WolfEnum interface {
	String() string
	Name() string
	Value() int
}

// wolfEnumBase is the internal, non-exported type
type wolfEnumBase struct{ value int }

// Value() returns the enum value
func (eb wolfEnumBase) Value() int { return eb.value }

// String() returns the name of the enum
func (eb wolfEnumBase) String() string { return eb.Name() }

// Name() is the enum name as you use it in Go code,
// needs to be overriden by inheriting types
func (eb wolfEnumBase) Name() string { return "" }

// Wolf1 is the enum type for 'piggy1' value
type Wolf1 struct{ wolfEnumBase }

// New is the constructor for a brand new WolfEnum with value 'piggy1'
func (eb Wolf1) New() WolfEnum { return Wolf1{wolfEnumBase{piggy1}} }

// String returns the enum name
func (eb Wolf1) String() string { return eb.Name() }

// Name returns always "Wolf1" for this enum type
func (eb Wolf1) Name() string { return "Wolf1" }

// Wolf2 is the enum type for 'piggy2' value
type Wolf2 struct{ wolfEnumBase }

// New is the constructor for a brand new WolfEnum with value 'piggy2'
func (eb Wolf2) New() WolfEnum { return Wolf2{wolfEnumBase{piggy2}} }

// String returns the enum name
func (eb Wolf2) String() string { return eb.Name() }

// Name returns always "Wolf2" for this enum type
func (eb Wolf2) Name() string { return "Wolf2" }

// Wolf3 is the enum type for 'piggy3' value
type Wolf3 struct{ wolfEnumBase }

// New is the constructor for a brand new WolfEnum with value 'piggy3'
func (eb Wolf3) New() WolfEnum { return Wolf3{wolfEnumBase{piggy3}} }

// String returns the enum name
func (eb Wolf3) String() string { return eb.Name() }

// Name returns always "Wolf3" for this enum type
func (eb Wolf3) Name() string { return "Wolf3" }

// Wolf4 is the enum type for 'piggy4' value
type Wolf4 struct{ wolfEnumBase }

// New is the constructor for a brand new WolfEnum with value 'piggy4'
func (eb Wolf4) New() WolfEnum { return Wolf4{wolfEnumBase{piggy4}} }

// String returns the enum name
func (eb Wolf4) String() string { return eb.Name() }

// Name returns always "Wolf4" for this enum type
func (eb Wolf4) Name() string { return "Wolf4" }

var internalWolfEnumValues = []WolfEnum{
	Wolf1{}.New(),
	Wolf2{}.New(),
	Wolf3{}.New(),
	Wolf4{}.New(),
}

// WolfEnumValues will return all available enum value types
func WolfEnumValues() []WolfEnum { return internalWolfEnumValues[:] }

// NewWolfFromValue will generate a valid enum from a value, or return nil in case of invalid value
func NewWolfFromValue(v int) (result WolfEnum) {
	switch v {
	case piggy1:
		result = Wolf1{}.New()
	case piggy2:
		result = Wolf2{}.New()
	case piggy3:
		result = Wolf3{}.New()
	case piggy4:
		result = Wolf4{}.New()
	}
	return
}
