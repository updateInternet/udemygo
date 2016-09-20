// returns half of integer and if the int is odd or even

package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(45)
	fmt.Println(h, even)
}
