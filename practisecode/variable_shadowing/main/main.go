package main

import "fmt"

// Max returns integer, visible outside of package
func Max(x int) int {
	return 42 + x
}

func main() {
	max := Max(7)
	fmt.Println(max) // max is now the result, not the fucntion
	Max(9)
}

// dont do this; bad coding practise
