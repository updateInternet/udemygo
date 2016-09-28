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
	buckets := make([][]string, 12)
	// Create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		// 12 buckets
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, "-", len(buckets[i]))
	}
	// Print the words in one of the buckets
	fmt.Println(buckets[6])
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

	// more uneven distribution
	//return len(word) * buckets
}
