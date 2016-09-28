package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type Doublezero struct {
	Person
	code string
	kill bool
}

func main() {
	p1 := Doublezero{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   45,
		},
		code: "007",
		kill: true,
	}
	fmt.Println(p1.Person, p1.code, p1.kill)
}
