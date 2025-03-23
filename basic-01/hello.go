// •	Every Go program starts with a package declaration.
// •	package main means this is an executable program (Go will look for func main() to start execution).
// •	If it was package utils, it would be a library (reusable code, no direct execution).

package main

import "fmt"

//	import lets you use code from other packages
//
// fmt is a standard Go package for formatting text (printing to console, reading input, etc.).
func main() {
	fmt.Println("Hello world, my first go code")
}
