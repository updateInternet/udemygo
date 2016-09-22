package main

import "fmt"

func main() {

	greet := map[int]string{
		0: "Good morning",
		1: "How are you",
		2: "AM fine",
		3: "See you later",
	}

	fmt.Println(greet) // shouldnt be in order
	delete(greet, 3)

	if val, exists := greet[3]; exists {
		fmt.Println("This value exists", "val: ", val, "exists: ", exists)
	} else {
		fmt.Println("This value doesnt exist", "val: ", val, "exists: ", exists)
	}
}
