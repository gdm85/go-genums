package main

// *** generated with go-genums ***

// DayEnum is the the enum interface that can be used
type DayEnum interface {
	String() string
	Value() dayEnum
	uniqueDayMethod()
}

// dayEnumBase is the internal, non-exported type
type dayEnumBase struct{ value dayEnum }

// Value() returns the enum value
func (eb dayEnumBase) Value() dayEnum { return eb.value }

// String() returns the enum name as you use it in Go code,
// needs to be overriden by inheriting types
func (eb dayEnumBase) String() string { return "" }

// Sunday is the enum type for 'daySunday' value
type Sunday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'daySunday'
func (Sunday) New() DayEnum { return Sunday{dayEnumBase{daySunday}} }

// String returns always "Sunday" for this enum type
func (Sunday) String() string { return "Sunday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Sunday) uniqueDayMethod() {}

// Monday is the enum type for 'dayMonday' value
type Monday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'dayMonday'
func (Monday) New() DayEnum { return Monday{dayEnumBase{dayMonday}} }

// String returns always "Monday" for this enum type
func (Monday) String() string { return "Monday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Monday) uniqueDayMethod() {}

// Tuesday is the enum type for 'dayTuesday' value
type Tuesday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'dayTuesday'
func (Tuesday) New() DayEnum { return Tuesday{dayEnumBase{dayTuesday}} }

// String returns always "Tuesday" for this enum type
func (Tuesday) String() string { return "Tuesday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Tuesday) uniqueDayMethod() {}

// Wednesday is the enum type for 'dayWednesday' value
type Wednesday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'dayWednesday'
func (Wednesday) New() DayEnum { return Wednesday{dayEnumBase{dayWednesday}} }

// String returns always "Wednesday" for this enum type
func (Wednesday) String() string { return "Wednesday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Wednesday) uniqueDayMethod() {}

// Thursday is the enum type for 'dayThursday' value
type Thursday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'dayThursday'
func (Thursday) New() DayEnum { return Thursday{dayEnumBase{dayThursday}} }

// String returns always "Thursday" for this enum type
func (Thursday) String() string { return "Thursday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Thursday) uniqueDayMethod() {}

// Friday is the enum type for 'dayFriday' value
type Friday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'dayFriday'
func (Friday) New() DayEnum { return Friday{dayEnumBase{dayFriday}} }

// String returns always "Friday" for this enum type
func (Friday) String() string { return "Friday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Friday) uniqueDayMethod() {}

// Saturday is the enum type for 'daySaturday' value
type Saturday struct{ dayEnumBase }

// New is the constructor for a brand new DayEnum with value 'daySaturday'
func (Saturday) New() DayEnum { return Saturday{dayEnumBase{daySaturday}} }

// String returns always "Saturday" for this enum type
func (Saturday) String() string { return "Saturday" }

// uniqueDayMethod() guarantees that the enum interface cannot be mis-assigned with others defined with an otherwise identical signature
func (Saturday) uniqueDayMethod() {}

var internalDayEnumValues = []DayEnum{
	Sunday{}.New(),
	Monday{}.New(),
	Tuesday{}.New(),
	Wednesday{}.New(),
	Thursday{}.New(),
	Friday{}.New(),
	Saturday{}.New(),
}

// DayEnumValues will return a slice of all allowed enum value types
func DayEnumValues() []DayEnum { return internalDayEnumValues[:] }

// NewDayFromValue will generate a valid enum from a value, or return nil in case of invalid value
func NewDayFromValue(v dayEnum) (result DayEnum) {
	switch v {
	case daySunday:
		result = Sunday{}.New()
	case dayMonday:
		result = Monday{}.New()
	case dayTuesday:
		result = Tuesday{}.New()
	case dayWednesday:
		result = Wednesday{}.New()
	case dayThursday:
		result = Thursday{}.New()
	case dayFriday:
		result = Friday{}.New()
	case daySaturday:
		result = Saturday{}.New()
	}
	return
}

// MustGetDayFromValue is the same as NewDayFromValue, but will panic in case of conversion failure
func MustGetDayFromValue(v dayEnum) DayEnum {
	result := NewDayFromValue(v)
	if result == nil {
		panic("invalid DayEnum value cast")
	}
	return result
}
