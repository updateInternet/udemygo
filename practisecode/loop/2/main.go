package main

import "fmt"

/* func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
*/

/* FOR with no condition - will run forever
func main() {
i:=0
for {
fmt.Println(i)
}
}
*/

/* func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
*/

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 10 {
			break
		}
	}
}
