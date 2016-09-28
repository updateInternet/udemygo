package main

import "fmt"

func main() {
	n := HashBucket("Hey how are you", 15)
	fmt.Println(n)
}

// HashBucket rune and int32 are the same
func HashBucket(word string, buckets rune) int32 {
	letter := rune(word[7])
	// letter := int(word[7])
	fmt.Println(letter)
	fmt.Println(buckets)
	bucket := letter % buckets
	return bucket
}
