package main

import "fmt"

// we declare the type person
type person struct {
	first string // these are fields
	last  string
	age   int
	// type can also have tags - used mostly for json related
}

func main() {
	p1 := person{"R", "s", 26} // this is called struct literal
	p2 := person{"Mr", "V", 30}
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p1.first, p1.last, p1.age)
}
