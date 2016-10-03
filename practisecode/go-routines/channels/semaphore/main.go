package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}() // anonymous self executing functions

	go func() {
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	// wrong way to do it
	// <-done
	// <-done
	// close(c)

	for n := range c {
		fmt.Println(n)
	}
}
