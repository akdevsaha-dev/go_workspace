package main

import "fmt"

func main() {
	const day = 2
	switch day {
	case 1:
		fmt.Print("monday")
	case 2:
		fmt.Print("tuesday")
	case 3:
		fmt.Print("wednesday")
	default:
		fmt.Print("holiday")
	}
}
