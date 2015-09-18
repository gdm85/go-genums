package main

//go:generate sh -c "go run ../../genums.go Day day dayEnum enum.go > generatedEnum.go"

type dayEnum int

// this is the user-provided canonical Go-style enum made up of constants
// notice how it is separated from the main.go file which references the generated
// enum declarations
const (
	dayMonday dayEnum = 1 + iota
	dayTuesday
	dayWednesday
	dayThursday
	dayFriday
	daySaturday
	daySunday
)
