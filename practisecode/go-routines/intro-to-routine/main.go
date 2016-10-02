package main

import "fmt"

func main() {
	go foo() // this will result in abnominal results
	bar()
}

func foo() {
	for i := 100; i < 106; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for b := 0; b < 15; b++ {
		fmt.Println("Bar:", b)
	}
}
