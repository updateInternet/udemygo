package main

import (
	"fmt"
	"sort"
)

type superhero []string

func main() {
	s := superhero{"SuperMan", "BatMan", "X-Men", "Deadpool"}
	sort.Strings(s)
	fmt.Println(s)
}
