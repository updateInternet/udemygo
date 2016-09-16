package main

import "fmt"

const (
	_  = iota             // ignore 0
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {
	fmt.Printf("%d %d \n", kb, mb)
	fmt.Printf("%b %b \n", kb, mb)
}
