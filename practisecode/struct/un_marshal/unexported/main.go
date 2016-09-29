package main

import (
	"encoding/json"
	"fmt"
)

// Person has all non-exported fields
type Person struct {
	first       string
	last        string
	age         int
	notexported int
}

func main() {
	p1 := Person{"James", "Bond", 48, 56}
	fmt.Println(p1)
	// since fileds are not exported, below will be empty
	// fields have to be exported to go outside of package
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
