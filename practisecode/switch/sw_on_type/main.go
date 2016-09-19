package main

/* Switch on types
-- normally we switch on value or variables as seen before
-- go allowes to switch on type of variable
*/

import "fmt"

// Contact type
type Contact struct {
	greeting string
	name     string
}

// SwitchOnType we will learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting "x is of this type"
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("this is of type Contact")
	default:
		fmt.Println("unkown type")
	}
}

func main() {
	SwitchOnType(1086)
	SwitchOnType("YOLO")
	var t = Contact{"Good to see you", "V"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
