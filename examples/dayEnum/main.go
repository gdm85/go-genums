package main

import "fmt"

func main() {
	fmt.Println("Days of the week:")
	for _, day := range DayEnumValues() {
		fmt.Println(day)
	}

	day := NewDayFromValue(1)
	if day == nil {
		panic("could not create Tuesday!")
	}
	fmt.Println("Tuesday is", day)

	if day != (Tuesday{}).New() {
		panic("Tuesday it's not Tuesday!")
	}

	day = NewDayFromValue(100)
	if day != nil {
		panic("imaginary day created!")
	}

	// example of safer evaluation with a type switch
	// loop will pick only odd Mondays
	for _, day := range DayEnumValues() {
		switch day.(type) {
		case Monday:
			fmt.Println("Found a Monday:", day)
		default:
			// ignore all other days
		}
	}
}
