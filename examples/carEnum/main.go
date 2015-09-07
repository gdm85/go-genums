package main

import "fmt"

func main() {
	ferrari := Ferrari{}.New()
	aston := NewCarFromValue(carType{"Aston Martin", "Lagonda", 6})
	lamborghini := Lamborghini{}.New()

	fmt.Println("Ferrari is", ferrari)
	fmt.Println("Aston Martin is", aston)
	fmt.Println("Lamborghini is", lamborghini)
}
