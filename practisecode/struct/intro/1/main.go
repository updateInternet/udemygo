package main

import "fmt"

// this is almost like foo is an alias for int
type foo int // this is not idiomatic

func main() {
	var myAge foo
	myAge = 30
	// prints type of myAge
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 44
	// fmt.Println(myAge + yourAge) //mismatched types foo and int will not work

	// this convertion works
	fmt.Println(int(myAge) + yourAge)
}
