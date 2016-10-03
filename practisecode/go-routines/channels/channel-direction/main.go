package main

/* The optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.

chan T          can be used to send and receive values of type T
chan<- float64  can only be used to send float64s
<-chan int      can only be used to receive ints

https://golang.org/ref/spec#Channel_types
*/

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
