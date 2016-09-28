package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//get moby dick book
	res, err := http.Get("http://fastss.csg.uzh.ch/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts {length and capacity of 200}
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	// 65:123 which is A to Z
	fmt.Println(buckets[65:123])
}

// HashBucket returns string alphabets in int buckets
func HashBucket(word string) int {
	return int(word[0])
}
