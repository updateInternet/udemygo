package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// wg.Add(1) -> this will introduce a race-condition,
		// moving this to global variable will solve the race condition problem
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}() // anonymous self executing functions

	go func() {
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait() // waiting for wg to finish
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
