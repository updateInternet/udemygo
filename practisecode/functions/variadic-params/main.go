package main

import "fmt"

func average(numbers ...float64) float64 {
	fmt.Println(numbers)
	total := 0.0 // or var total float64 (which is initalize 0 by default)
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	n := average(4, 8, 10, 24)
	fmt.Println(n)
}
