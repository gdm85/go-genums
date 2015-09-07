package main

//go:generate sh -c "go run ../../genums.go Day day int enum.go > generatedEnum.go"

// this is the user-provided canonical Go-style enum made up of constants
// notice how it is separated from the main.go file which references the generated
// enum declarations
const (
	dayMonday = iota
	dayTuesday
	dayWednesday
	dayThursday
	dayFriday
	daySaturday
	daySunday
)
