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

	// create slice to hold counts {length and capacity of 12}
	buckets := make([]int, 12)
	// Loop over the words
	for scanner.Scan() { // Scan() is bool; if true keep looping
		// 12 buckets
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

// HashBucket returns string alphabets in int buckets
func HashBucket(word string, buckets int) int {
	/*	letter := int(word[0])
		bucket := letter % buckets
		return bucket
	*/
	// more even distribution of buckets
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
