package main

import "fmt"

func main() {
	greet("Howdy", "V")
}

func greet(say, something string) { // 2 params
	fmt.Println(say, something)
}

// greet is declared with paramter
// when calling greet, pass in a argument
