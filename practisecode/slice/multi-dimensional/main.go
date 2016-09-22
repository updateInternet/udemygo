package main

import "fmt"

func main() {
	trans := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		t := make([]int, 0) // or var t []int
		for j := 0; j < 5; j++ {
			t = append(t, j)
		}
		fmt.Println(t)
		trans = append(trans, t)
	}
	fmt.Println(trans)
}
