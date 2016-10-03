package main

import (
	"fmt"
	"time"
)

func main() {
	// this is how we create a channel
	// communicate int via channel
	// 	this is an unbuffered channel
	// 	buffered channel is (chan int, 10) --> takes 10 things
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i // we are putting a number (i) onto the channel
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}
