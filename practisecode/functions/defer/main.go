package main

import "fmt"

func hello() {
	fmt.Println("Hello everyone")
}

func mrng() {
	fmt.Println("Have a wonderful day ahead")
}

func main() {
	defer mrng() // defers this function right before main exists
	hello()
}
