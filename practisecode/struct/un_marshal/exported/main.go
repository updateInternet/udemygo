package main

import (
	"encoding/json"
	"fmt"
)

// Person has exported and non-exported fields
type Person struct {
	First       string
	Last        string
	Age         int
	notexported int
}

func main() {
	p1 := Person{"James", "Bond", 48, 56}
	fmt.Println(p1)
	// bytes of slice
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
