package main

import (
	"fmt"
	"sort"
)

type superhero []string

// need these methods for type Interface interface
func (s superhero) Len() int           { return len(s) }
func (s superhero) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s superhero) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	sh := superhero{"SuperMan", "BatMan", "X-Men", "Deadpool"}
	fmt.Println(sh)
	sort.Sort(sh) // takes an interface
	fmt.Println(sh)
}
