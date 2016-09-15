package main

import "fmt"

// returns a function
func wrapper() func() int {
	x := 100
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
