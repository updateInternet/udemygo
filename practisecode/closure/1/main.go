package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x) // anything from outer scope is accessible in inner scope
		y := "This is from inner closure"
		fmt.Println(y)
	} // Inner scope
	// fmt.Println(y)
	// Outside scope of y; inner scope not accessible to outer scope
} // Outer scope
