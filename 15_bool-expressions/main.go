package main

import "fmt"

func main() {
	ifRight("Right")
}

func ifRight (Eric string) bool {
	if (Eric == "Right") {
		fmt.Println (Eric)
		return true
	}
	return false
}