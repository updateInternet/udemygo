package main

import "fmt"

func greeting() func() string {
	return func() string {
		return "Hello"
	}
}

func main() {
	greet := greeting()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet)
}
