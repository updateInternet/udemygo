package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = `Big O notation is a mathematical notation that describes the
	limiting behavior of a function when the argument tends towards a particular
	value or infinity. It is a 	member of a family of notations invented by Paul
	Bachmann,[1] Edmund Landau,[2] and others, collectively called
	Bachmann-Landau notation or asymptotic notation.`

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
