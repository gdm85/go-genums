package main

import "fmt"

func main() {
	for _, wolf := range WolfEnumValues() {
		fmt.Println("Enum value: ", wolf)
	}

	ev1 := Wolf1{}.New()
	ev2 := Wolf2{}.New()

	if ev1 != ev2 {
		fmt.Println(ev1, " differs from ", ev2)
	} else {
		fmt.Println(ev1, " is the same as ", ev2)
	}

	ev1 = Wolf3{}.New()
	ev2 = Wolf3{}.New()

	if ev1 != ev2 {
		fmt.Println(ev1, " differs from ", ev2)
	} else {
		fmt.Println(ev1, " is the same as ", ev2)
	}

	ev1 = NewWolfFromValue(1)
	if ev1 == nil {
		panic("could not create wolf!")
	}
	ev2 = Wolf2{}.New()

	if ev1 != ev2 {
		fmt.Println(ev1, " differs from ", ev2)
	} else {
		fmt.Println(ev1, " is the same as ", ev2)
	}

	ev1 = NewWolfFromValue(100)
	if ev1 != nil {
		panic("imaginary wolf created!")
	}

	// example of safer evaluation with a type switch
	// loop will pick only odd wolfs
	for _, wolf := range WolfEnumValues() {
		switch wolf.(type) {
		case Wolf1, Wolf3:
			fmt.Println("Found an odd wolf:", wolf)
		default:
			// ignore all other wolves
		}
	}
}
