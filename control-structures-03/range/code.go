package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}
}
