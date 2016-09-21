package main

import "fmt"

func main() {
	slices := []string{"Yo,", "you", "only", "live", "once"}
	fmt.Println(slices)
	fmt.Println(slices[1:3]) // slicing a slice
	fmt.Println(slices[4])   // imdex access
}
