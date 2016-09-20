package main

import "fmt"

func main() {
	x := 46
	half := func() (int, bool) {
		return x / 2, x%2 == 0
	}
	fmt.Println(half())
}

// use func expression to contain the scope of x
