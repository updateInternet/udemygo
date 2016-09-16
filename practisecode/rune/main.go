package main

import "fmt"

func main() {
	for i := 120; i <= 125; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	foo := 'a' // rune
	boo := "b" // string
	fmt.Println(foo, boo)
	fmt.Printf("%T  %T \n", foo, boo)
	fmt.Printf("%T \n", rune(foo))
	// fmt.Printf("%T \n", rune(boo)) cannot convert boo type string to
	// type rune (int32)
}
