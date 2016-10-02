package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"SuperMan", "BatMan", "X-Men", "Deadpool"}
	// sort.Strings(s)
	fmt.Println(s)
	//sort.StringSlice(s).Sort() // StringSlice is an interface
	//fmt.Println(s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	// reverse of string
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
