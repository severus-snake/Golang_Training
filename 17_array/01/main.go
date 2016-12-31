package main

import "fmt"

func main() {
	var x = [3]int{58, 101, 12} //literals defined for the array require the declaration operator
	var y [5]int //declaration operator not needed if you're not defining the objects in your array
	fmt.Println(x,y)
	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(x[2],y[4])
	x[2] = 777
	y[4] = 24
	fmt.Println(x[2],y[4])
}
