package main

import "fmt"

func main() {
	switch "V" {
	case "A":
		fmt.Println("Wassup A")
	case "B":
		fmt.Println("Wassup B")
	case "V", "H":
		fmt.Println("Yo V, err H")
		fallthrough
	case "Z":
		fmt.Println("End")
		fallthrough
	case "T":
		fmt.Println("TTT")
	case "F":
		fmt.Println("FFF")
	}
}
