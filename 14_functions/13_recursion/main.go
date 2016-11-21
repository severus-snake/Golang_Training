package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}

// 4 * 3 * 2 * 1 * 1 = 24 (that's what happening in this recursion function)

// Recursion is when a function calls itself