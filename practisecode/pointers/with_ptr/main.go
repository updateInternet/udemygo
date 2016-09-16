package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) // memory address in func zero
	*z = 0
}

func main() {
	x := 105
	fmt.Println(&x) // different way to print memory address
	zero(&x)
	fmt.Println(x) // x is 0 now using pointers
}
