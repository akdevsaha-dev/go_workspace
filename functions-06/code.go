package main

import "fmt"

// return sigle value
func add(a int, b int) int {
	return a + b
}

// return multiple value
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// variadic function
// use ... to accept multiple arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total = total + num
	}
	return total
}


func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)
	ans := add(10, 20)
	fmt.Println(ans)
	fmt.Println(sum(5, 5, 10, 10, 50))
}
