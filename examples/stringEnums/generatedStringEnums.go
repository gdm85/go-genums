package main

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

// ColorYellow is the enum type for 'colourYellow' value
type ColorYellow struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'colourYellow'
func (ColorYellow) New() ColorEnum { return ColorYellow{colorEnumBase{colourYellow}} }

// String returns always "ColorYellow" for this enum type
func (ColorYellow) String() string { return "ColorYellow" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (ColorYellow) uniqueColorMethod() {}

// ColorRed is the enum type for 'colourRed' value
type ColorRed struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'colourRed'
func (ColorRed) New() ColorEnum { return ColorRed{colorEnumBase{colourRed}} }

// String returns always "ColorRed" for this enum type
func (ColorRed) String() string { return "ColorRed" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (ColorRed) uniqueColorMethod() {}

// ColorBrown is the enum type for 'colourBrown' value
type ColorBrown struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'colourBrown'
func (ColorBrown) New() ColorEnum { return ColorBrown{colorEnumBase{colourBrown}} }

// String returns always "ColorBrown" for this enum type
func (ColorBrown) String() string { return "ColorBrown" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (ColorBrown) uniqueColorMethod() {}

// ColorGreen is the enum type for 'colourGreen' value
type ColorGreen struct{ colorEnumBase }

// New is the constructor for a brand new ColorEnum with value 'colourGreen'
func (ColorGreen) New() ColorEnum { return ColorGreen{colorEnumBase{colourGreen}} }

// String returns always "ColorGreen" for this enum type
func (ColorGreen) String() string { return "ColorGreen" }

// uniqueColorMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (ColorGreen) uniqueColorMethod() {}

var internalColorEnumValues = []ColorEnum{
	ColorYellow{}.New(),
	ColorRed{}.New(),
	ColorBrown{}.New(),
	ColorGreen{}.New(),
}

// ColorEnumValues will return all available enum value types
func ColorEnumValues() []ColorEnum { return internalColorEnumValues[:] }

// NewColorFromValue will generate a valid enum from a value, or return nil in case of invalid value
func NewColorFromValue(v string) (result ColorEnum) {
	switch v {
	case colourYellow:
		result = ColorYellow{}.New()
	case colourRed:
		result = ColorRed{}.New()
	case colourBrown:
		result = ColorBrown{}.New()
	case colourGreen:
		result = ColorGreen{}.New()
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
