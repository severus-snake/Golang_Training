package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}

// A slice is a reference type, if you change a value of a reference type
// it will also change the value at it's memory address