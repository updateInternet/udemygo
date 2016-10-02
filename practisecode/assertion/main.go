package main

import "fmt"

func main() {
	// this will result in assertion error
	/*	name := "IronMan"
		str, ok := name.(string) // this is not a interface type
	*/
	var name interface{} = "IronMan"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}

	var val interface{} = 72
	// val + 6 (mismatched types interface {} and int)
	// fmt.Println(val + 6)
	fmt.Println(val.(int) + 6) // if you assert val as int, it wil work

	rem := 65.32
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem)) // casting

}
