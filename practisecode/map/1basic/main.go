package main

import "fmt"

func main() {

	var greet = map[string]string{} // composite literal
	//or var greet = make(map[string]string)
	greet["V"] = "Hello"
	greet["R"] = "Bonjour"
	fmt.Println(greet)
	fmt.Println(greet == nil)

	mygreet := map[string]string{ // shorthand
		"V": "Hello",
		"R": "Bonjour",
	} // composite literal
	mygreet["S"] = "Howdy" // adding entry
	fmt.Println(mygreet)
	mygreet["R"] = "How are ya"
	fmt.Println(mygreet)
	delete(mygreet, "S")
	fmt.Println(mygreet)
	fmt.Println(len(mygreet))
}

/*
var greet map[string]string
greet["V"] = "Hello"
greet["R"] = "Bonjour"
fmt.Println(greet)
fmt.Println(greet == nil)

This will fail due to no initalization of map

var greet map[string]string
// greet["V"] = "Hello"
// greet["R"] = "Bonjour"
fmt.Println(greet)
fmt.Println(greet == nil)

This will return True
*/
