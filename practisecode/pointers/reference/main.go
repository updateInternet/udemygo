package main

import "fmt"

func main() {

	a := 55

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // referencing a memory address
	fmt.Println(b)
	fmt.Println(*b) // deferencing a memory address - 55
}

// b a pointer to the memory address
// b is of type "int pointer"
