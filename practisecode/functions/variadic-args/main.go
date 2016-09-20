package main

import "fmt"

func average(numbers ...float64) float64 {
	var total float64
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	data := []float64{8, 10, 24, 32, 46, 54}
	fmt.Println(data)

	n := average(data...)
	fmt.Println(n)
}
