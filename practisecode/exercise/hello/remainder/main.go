package main

import "fmt"

var x, y int

func main() {
	fmt.Print("Enter the smallest number in your mind:")
	fmt.Scan(&x)
	fmt.Print("Enter the largest number in your mind:")
	fmt.Scan(&y)
	fmt.Println("remainder:", y/x)
}
