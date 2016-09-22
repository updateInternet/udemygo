package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 2)
	// 0 is length - number of elements referred to by the slice
	// 2 is capacity - number of elements in the underlying array
	fmt.Println("-------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-------------")

	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, i)
		// appending beyond the capacity(2) to a slice
		fmt.Println("Len:", len(mySlice), "value:", i, "Capacity:", cap(mySlice))
	}
	myFirstSlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9, 0}
	// appending a slice to a slice
	myFirstSlice = append(myFirstSlice, myOtherSlice...) // notice the syntax

	fmt.Println(myFirstSlice)
}
