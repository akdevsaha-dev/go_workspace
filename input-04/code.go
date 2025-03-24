package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	println("Please enter you name")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is", name)
}
