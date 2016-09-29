package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James Bond", 38}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
}
