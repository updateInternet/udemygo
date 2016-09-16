package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z)
	z = 0
}

func main() {
	x := 105
	fmt.Printf("%p\n", &x) // memory address in main
	fmt.Println(&x)        // different way to print memory address
	zero(x)
	fmt.Println(x) // x is still 105
}
