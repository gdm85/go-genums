package main

// *** generated with go-genums ***

// ColorEnum is the the enum interface that can be used
type ColorEnum interface {
	String() string
	Value() string
	uniqueColorMethod()
}

// colorEnumBase is the internal, non-exported type
type colorEnumBase struct{ value string }

// Value() returns the enum value
func (eb colorEnumBase) Value() string { return eb.value }

// String() returns the enum name as you use it in Go code,
// needs to be overriden by inheriting types
func (eb colorEnumBase) String() string { return "" }

// Green is the enum type for 'clGreen' value
type Green struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clGreen'
func (Green) New() ColorEnum { return Green{colorEnumBase{clGreen}} }

// String returns always "Green" for this enum type
func (Green) String() string { return "Green" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Green) uniqueColorMethod() {}

// Brown is the enum type for 'clBrown' value
type Brown struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clBrown'
func (Brown) New() ColorEnum { return Brown{colorEnumBase{clBrown}} }

// String returns always "Brown" for this enum type
func (Brown) String() string { return "Brown" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Brown) uniqueColorMethod() {}

// Pink is the enum type for 'clPink' value
type Pink struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clPink'
func (Pink) New() ColorEnum { return Pink{colorEnumBase{clPink}} }

// String returns always "Pink" for this enum type
func (Pink) String() string { return "Pink" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Pink) uniqueColorMethod() {}

// White is the enum type for 'clWhite' value
type White struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clWhite'
func (White) New() ColorEnum { return White{colorEnumBase{clWhite}} }

// String returns always "White" for this enum type
func (White) String() string { return "White" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (White) uniqueColorMethod() {}

// Black is the enum type for 'clBlack' value
type Black struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clBlack'
func (Black) New() ColorEnum { return Black{colorEnumBase{clBlack}} }

// String returns always "Black" for this enum type
func (Black) String() string { return "Black" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Black) uniqueColorMethod() {}

// Red is the enum type for 'clRed' value
type Red struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clRed'
func (Red) New() ColorEnum { return Red{colorEnumBase{clRed}} }

// String returns always "Red" for this enum type
func (Red) String() string { return "Red" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Red) uniqueColorMethod() {}

// Yellow is the enum type for 'clYellow' value
type Yellow struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'clYellow'
func (Yellow) New() ColorEnum { return Yellow{colorEnumBase{clYellow}} }

// String returns always "Yellow" for this enum type
func (Yellow) String() string { return "Yellow" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Yellow) uniqueColorMethod() {}

var internalColorEnumValues = []ColorEnum{
	Green{}.New(),
	Brown{}.New(),
	Pink{}.New(),
	White{}.New(),
	Black{}.New(),
	Red{}.New(),
	Yellow{}.New(),
}

// ColorEnumValues will return a slice of all allowed enum value types
func ColorEnumValues() []ColorEnum { return internalColorEnumValues[:] }

// NewColorFromValue will generate a valid enum from a value, or return nil in case of invalid value
func NewColorFromValue(v string) (result ColorEnum) {
	switch v {
	case clGreen:
		result = Green{}.New()
	case clBrown:
		result = Brown{}.New()
	case clPink:
		result = Pink{}.New()
	case clWhite:
		result = White{}.New()
	case clBlack:
		result = Black{}.New()
	case clRed:
		result = Red{}.New()
	case clYellow:
		result = Yellow{}.New()
	}
	return
}

// MustGetColorFromValue is the same as NewColorFromValue, but will panic in case of conversion failure
func MustGetColorFromValue(v string) ColorEnum {
	result := NewColorFromValue(v)
	if result == nil {
		panic("invalid ColorEnum value cast")
	}
	return result
}
