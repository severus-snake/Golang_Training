package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue //continue makes the loop go back to the top of the loop
			// and ignores all the lines of the loop code below it
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
