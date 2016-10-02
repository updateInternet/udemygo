package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 12
	var b = 14.975
	fmt.Println(float64(a) + b) // this is widening; int --> float64
	fmt.Println(a + int(b))     // this is shortening; float64 --> int

	fmt.Println([]byte("Yoi")) // conversion : string to byte

	var c = "34" // string which is ascii underneath
	var d = 78
	// multivalue return, ignoring errors
	e, _ := strconv.Atoi(c) // conversion : ascii to int (atoi)
	fmt.Println(e + d)

	f := 2
	g := "Iam Ironman part:" + strconv.Itoa(f) // conversion: int to ascii
	fmt.Println(g)
}
