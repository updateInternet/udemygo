package main

import "fmt"

// Person is a type which is exported
// fields are not exported
type Person struct {
	first string
	last  string
	age   int
}

// Doublezero is also a type with type Person inherited
type Doublezero struct {
	Person
	code string
	kill bool
}

func (p Person) Greeting() {
	fmt.Println("My name is Bond; James bond")
}

func (dz Doublezero) Greeting() {
	fmt.Println("Iam Tony Stark; aka Ironman")
}

func main() {
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   45,
	}
	p2 := Doublezero{
		Person: Person{
			first: "Iron",
			last:  "Man",
			age:   37,
		},
		code: "stark",
		kill: false,
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
