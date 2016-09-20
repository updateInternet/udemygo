package main

import "fmt"

func greet(say, something string) (s string) { // 2 params
	s = fmt.Sprint(say, something)
	return
}

func main() {
	s := greet("Howdy ", "Vas")
	fmt.Println(s)
	// fmt.Println(greet("Howdy ", "Vas"))
}
