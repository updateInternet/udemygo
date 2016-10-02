package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{54, 654, 78, 32, 1, 23, 9854, 67}
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	sort.Ints(n)
	fmt.Println(n)
}
