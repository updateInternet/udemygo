// Slice with string
package main

import "fmt"

func main() {
	slices := []string{"Yo", "you", "only", "live", "once"}

	for i, entry := range slices {
		fmt.Println(i, entry) // i is the index of slice, entry is the value
	}

	for j := 0; j < len(slices); j++ {
		fmt.Println(slices[j])
	}

	fmt.Println(slices)
	fmt.Println(slices[1:3]) // slicing a slice
	fmt.Println(slices[4])   // imdex access
	fmt.Println(slices[2:])
	fmt.Println(slices[:2])
}
