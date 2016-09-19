package main

import "fmt"

var name string

func main() {
	fmt.Print("What is your name:")
	fmt.Scan(&name)
	fmt.Println("My name is", name)
}
