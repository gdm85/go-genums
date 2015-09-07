package main

import "fmt"

func main() {
	black := Black{}.New()
	green := NewColorFromValue("00FF00")
	white := White{}.New()

	fmt.Println("Black is", black)
	fmt.Println("Green is", green)
	fmt.Println("White is", white)
}
