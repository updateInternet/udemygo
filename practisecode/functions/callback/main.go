package main

import "fmt"

func visit(numbers []int, anony func(int)) {
	for _, n := range numbers {
		anony(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing a function
