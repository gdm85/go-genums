package main

//go:generate sh -c "go run ../../genums.go Color cl string enum.go > generatedEnum.go"

// this is the user-provided canonical Go-style enum made up of constants
// notice how it is separated from the main.go file which references the generated
// enum declarations
const (
	clBlack  = "000000"
	clRed    = "FF0000"
	clYellow = "FFFF00"
	clGreen  = "00FF00"
	clBrown  = "8B4513"
	clPink   = "F08080"
	clWhite  = "FFFFFF"
)
