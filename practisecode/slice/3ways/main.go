package main

import "fmt"

func main() {
	a := []string{} // underlying array doesnt have any spots
	as := [][]string{}
	fmt.Println(a)
	fmt.Println(as)
	fmt.Println(a == nil)

	var b []string // 0 value of slice which is nil
	var bs [][]string
	fmt.Println(b)
	fmt.Println(bs)
	fmt.Println(b == nil)

	// idiomatic way of representing slices
	c := make([]string, 5)
	cs := make([][]string, 7)
	fmt.Println(c)
	fmt.Println(cs)
	fmt.Println(c == nil)

	//Incrementing a slice
	d := make([]int, 6, 100) // Other way new([100]int)[0:6]
	fmt.Println(d[5])
	d[1] = 56
	fmt.Println(d[1])
	d[1]++
	// d[1] +=1 or d[1] = d[1] + 1
	fmt.Println(d[1])
}

// Slices are dynamic
