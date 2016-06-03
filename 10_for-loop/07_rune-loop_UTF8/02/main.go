package main

import "fmt"

func main() {
	for i := 50; i <= 202; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
