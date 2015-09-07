package main

//go:generate sh -c "go run ../../genums.go  Color colour string main.go > generatedStringEnums.go"

// this is the user-provided canonical Go-style enum made up of constants
const (
	colourYellow = "YELLOW"
	colourRed    = "RED"
	colourBrown  = "BROWN"
	colourGreen  = "GREEN"
)

func main() {
}
