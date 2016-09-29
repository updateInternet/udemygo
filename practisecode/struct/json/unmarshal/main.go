package main

import (
	"encoding/json"
	"fmt"
)

// Person is a struct
type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	// bytes of slice
	bs := []byte(`{"First": "James", "Last":"Bond", "wisdom score":20}`)
	json.Unmarshal(bs, &p1)
	fmt.Println("++++++++++++++++++++")
	fmt.Println(string(bs))
	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Printf("%T \n", p1)
}
