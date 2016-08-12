package main

import "fmt"

func main() {
	n := hashBucket("Go", 12) //what's going on here???
	fmt.Println(n)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	fmt.Println(letter)
	bucket := letter % buckets
	fmt.Println(bucket)
	return bucket
}
