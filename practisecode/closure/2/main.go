package main

import "fmt"

var x = 2 // this is always remembered

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
