# Golang User Input

In Go, you can take user input from the command line using the `fmt` package. Below is an example of reading simple input:


## IMP
fmt.Scan: Stops at thre first whitespace it encounters


## Using `fmt.Scanln` (Simple Input)
The `fmt.Scanln` function is useful for reading simple input like strings, integers, etc.

### Behavior of fmt.Scanln
fmt.Scanln reads input until it encounters a newline (\n), but it stops reading at the first whitespace when assigning values to variables.

If you provide more input than the number of variables, the extra input is ignored.

Example- if you enter john doe it will only print john


```go
package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // Reads input until a newline is encountered or a space is encountered
	fmt.Printf("Hello, %s!\n", name)
}
```


## Reading Multiple Inputs
You can also read multiple inputs using fmt.Scan or fmt.Scanln.

```go
package main

import (
	"fmt"
)

func main() {
	var firstName, lastName string
	fmt.Print("Enter your first name and last name: ")
	fmt.Scanln(&firstName, &lastName) // Reads two inputs separated by a space
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}
```


### Scanln vs bufio
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Using fmt.Scanln
	var name1 string
	fmt.Print("Enter your full name (fmt.Scanln): ")
	fmt.Scanln(&name1)
	fmt.Printf("fmt.Scanln: Hello, %s!\n", name1)

	// Using bufio.NewReader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name (bufio.NewReader): ")
	name2, _ := reader.ReadString('\n')
	fmt.Printf("bufio.NewReader: Hello, %s!\n", name2)
}
```

```bash
Enter your full name (fmt.Scanln): akdev saha
fmt.Scanln: Hello, akdev!
Enter your full name (bufio.NewReader): akdev saha
bufio.NewReader: Hello, akdev saha!
```