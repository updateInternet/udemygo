package main

import "fmt"

func main() {
	var x [50]string
	// var x []string --> this is a slice
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[40])
	for i := 65; i < 110; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[36])
}
