package main

import (
	"encoding/json"
	"fmt"
)

// Person has exported and non-exported fields
type Person struct {
	First string
	Last  string `json:"-"` // this doesnt send last field outside of package
	Age   int    `json:"wisdowm score"`
	// ^ replaces Age with ws whenexporting outside of package using tags
}

func main() {
	p1 := Person{"James", "Bond", 48}
	fmt.Println(p1)
	// bytes of slice
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
