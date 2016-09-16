package main

import "fmt"

func main() {

	a := 60
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // referencing a memory address
	fmt.Println(b)
	fmt.Println(*b) // deferencing a memory address - 55

	*b = 200 // The value at this address changes to 200 from 55
	fmt.Println(a)
}

// everything is PASS BY VALUE in go
// when we pass a memory address, we are passing a value
