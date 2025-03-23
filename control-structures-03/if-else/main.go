package main

import "fmt"

func main() {
	const age1 int = 30
	age2 := 25 // short hand declaration -- type inferred.
	if age1 > age2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
