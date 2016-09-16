package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello")) // converting hello into a slice of bytes
	for i := 100; i <= 200; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
		// using printf as below
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
