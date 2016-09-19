package main

import "fmt"

func main() {

	yep := "YOLO"

	switch { // without expression turns it to boolean
	case yep == "":
		fmt.Println("Wassup A")
	case len(yep) == 4:
		fmt.Println("YOLO YOLO")
		fallthrough // prints the next case too
	case yep == "V":
		fmt.Println("Yo V, err H")
	case yep == "Z":
		fmt.Println("End")
	}
}
