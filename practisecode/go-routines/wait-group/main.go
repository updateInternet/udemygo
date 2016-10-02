package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 100; i < 106; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for b := 0; b < 15; b++ {
		fmt.Println("Bar:", b)
	}
	wg.Done()
}
