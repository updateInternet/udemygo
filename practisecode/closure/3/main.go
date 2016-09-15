package main

import "fmt"

func main() {
	x := 10
	// assigning a function to variable increment
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/* anonymous function - no name

func expression - assinging a function to a variable
*/
