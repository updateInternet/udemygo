package main

// Note: Only the sender should close a channel, never the receiver.
// Sending on a closed channel will cause a panic.

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // we are putting a number (i) onto the channel
		}
		close(c) // close the channel, when channel is closed, no longer put values
		// on the channel
	}() // self-executes

	for n := range c {
		fmt.Println(n)
	}
}
