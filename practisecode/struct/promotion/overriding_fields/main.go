package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doublezero struct {
	person
	code string
	kill bool
}

func main() {
	p1 := doublezero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   45,
		},
		code: "007",
		kill: true,
	}
	p2 := doublezero{
		person: person{
			first: "Iron",
			last:  "Man",
			age:   37,
		},
		code: "stark",
		kill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.first, p1.person.first)
	fmt.Println(p2.person, p2.code, p2.kill)
	fmt.Println(p2.first, p2.person.last, p2.age)

}
