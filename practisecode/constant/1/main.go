package main

import "fmt"

const (
	_          = iota
	def        = iota
	r   string = "R stands for Robot"
	s          = 101
)

func main() {

	fmt.Println("s is - ", s, r)
	fmt.Println(def)
}
