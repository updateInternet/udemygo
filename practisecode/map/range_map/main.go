package main

import "fmt"

func main() {

	greet := map[int]string{
		0: "Good morning",
		1: "How are you",
		2: "AM fine",
		3: "See you later",
	}

	// range loop over a map
	for k, v := range greet {
		fmt.Println(k, "=", v)
	}
}
