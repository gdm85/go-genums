package main

// *** generated with go-genums ***

// CarEnum is the the enum interface that can be used
type CarEnum interface {
	String() string
	Value() carType
	uniqueCarMethod()
}

// carEnumBase is the internal, non-exported type
type carEnumBase struct{ value carType }

// Value() returns the enum value
func (eb carEnumBase) Value() carType { return eb.value }

// String() returns the enum name as you use it in Go code,
// needs to be overriden by inheriting types
func (eb carEnumBase) String() string { return "" }

// Ferrari is the enum type for 'cFerrari' value
type Ferrari struct{ carEnumBase }

// New is the constructor for a brand new CarEnum with value 'cFerrari'
func (Ferrari) New() CarEnum { return Ferrari{carEnumBase{cFerrari}} }

// String returns always "Ferrari" for this enum type
func (Ferrari) String() string { return "Ferrari" }

// uniqueCarMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Ferrari) uniqueCarMethod() {}

// AstonMartin is the enum type for 'cAstonMartin' value
type AstonMartin struct{ carEnumBase }

// New is the constructor for a brand new CarEnum with value 'cAstonMartin'
func (AstonMartin) New() CarEnum { return AstonMartin{carEnumBase{cAstonMartin}} }

// String returns always "AstonMartin" for this enum type
func (AstonMartin) String() string { return "AstonMartin" }

// uniqueCarMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (AstonMartin) uniqueCarMethod() {}

// Lamborghini is the enum type for 'cLamborghini' value
type Lamborghini struct{ carEnumBase }

// New is the constructor for a brand new CarEnum with value 'cLamborghini'
func (Lamborghini) New() CarEnum { return Lamborghini{carEnumBase{cLamborghini}} }

// String returns always "Lamborghini" for this enum type
func (Lamborghini) String() string { return "Lamborghini" }

// uniqueCarMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Lamborghini) uniqueCarMethod() {}

// DeLorean is the enum type for 'cDeLorean' value
type DeLorean struct{ carEnumBase }

// New is the constructor for a brand new CarEnum with value 'cDeLorean'
func (DeLorean) New() CarEnum { return DeLorean{carEnumBase{cDeLorean}} }

// String returns always "DeLorean" for this enum type
func (DeLorean) String() string { return "DeLorean" }

// uniqueCarMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (DeLorean) uniqueCarMethod() {}

var internalCarEnumValues = []CarEnum{
	Ferrari{}.New(),
	AstonMartin{}.New(),
	Lamborghini{}.New(),
	DeLorean{}.New(),
}

// CarEnumValues will return a slice of all allowed enum value types
func CarEnumValues() []CarEnum { return internalCarEnumValues[:] }

// NewCarFromValue will generate a valid enum from a value, or return nil in case of invalid value
func NewCarFromValue(v carType) (result CarEnum) {
	switch v {
	case cFerrari:
		result = Ferrari{}.New()
	case cAstonMartin:
		result = AstonMartin{}.New()
	case cLamborghini:
		result = Lamborghini{}.New()
	case cDeLorean:
		result = DeLorean{}.New()
	}
	return
}

// MustGetCarFromValue is the same as NewCarFromValue, but will panic in case of conversion failure
func MustGetCarFromValue(v carType) CarEnum {
	result := NewCarFromValue(v)
	if result == nil {
		panic("invalid CarEnum value cast")
	}
	return result
}
