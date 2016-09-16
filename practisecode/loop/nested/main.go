package main

import "fmt"

func main() {
	// for initializer, condition, post
	for i := 0; i <= 2; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, "-", j)
		}
	}
}
