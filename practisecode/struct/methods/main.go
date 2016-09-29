package main

import "fmt"

// we declare the type person
type person struct {
	first string // these are fields
	last  string
	age   int
}

// (p person) is the receiver
// fullName() is the function name
// it returns string
func (p person) fullName() string {
	// any value of type person can now call method fullName()
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 48}
	p2 := person{"Iron", "Man", 37}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
