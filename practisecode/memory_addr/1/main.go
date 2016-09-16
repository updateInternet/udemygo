package main

import "fmt"

const meterstokms float64 = 0.001

var meters float64

func main() {
	fmt.Print("Enter meters ran: ")
	fmt.Scan(&meters)
	kms := meters * meterstokms
	fmt.Println(meters, "meters in", kms, "kms")
	fmt.Println("Memory address of kms:", &kms)
	fmt.Printf("Decimal value of memory address - %d \n", &kms)
}
