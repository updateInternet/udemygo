package main

import "fmt"

func main() {
	m := make(map[string]int) // map[key type]element/value type
	m["k1"] = 7
	m["k2"] = 20
	fmt.Println("map:", m)

	delete(m, "k1")
	fmt.Println("map:", m)

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, v)

	x, present := m["k2"]
	fmt.Println("ok?:", present, x)
}
