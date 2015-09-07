package main

//go:generate sh -c "go run ../../genums.go Car c carType enum.go > generatedEnum.go"

type carType struct {
	Brand, Model string
	Gears        int
}

// this is a more advanced enum that is based on variables instead of constants,
// and uses a more complex type (struct)
// notice how it is separated from the main.go file which references the generated
// enum declarations
var (
	cFerrari     = carType{"Ferrari", "Enzo", 6}
	cAstonMartin = carType{"Aston Martin", "Lagonda", 6}
	cLamborghini = carType{"Lamborghini", "Gallardo", 6}
	cDeLorean    = carType{"De Lorean", "DMC-12", 5}
)
