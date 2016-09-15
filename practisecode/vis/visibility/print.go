package visibility

import "fmt" // imports are only file level scope

// PrintVar prints my name and your name
func PrintVar() {
	fmt.Println(MyName)
	// yourName not exported but visibile within package level
	// this is package level scope
	// this is also like getter, not direct access but you can take it and show
	// value
	fmt.Println(yourName)
}
