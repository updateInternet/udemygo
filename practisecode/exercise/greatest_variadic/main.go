// Find the largest number using variadic
package main

import "fmt"

func main() {
	greatest := numbers(123, 12678, 4567, 23, 9)
	fmt.Println(greatest)
}

func numbers(n ...int) int {
	fmt.Printf("%T \n", n)
	var largest int
	for _, value := range n {
		if value > largest {
			largest = value
		}
	}
	return largest
}
