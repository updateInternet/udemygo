package main

import (
	"encoding/json"
	"os"
)

// Person is a struct
type Person struct {
	First    string
	Last     string
	Age      int
	noexport int
}

func main() {
	p1 := Person{"James", "Bond", 48, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
