package main

import "fmt"

func main() {

	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}

}

//Exclamation point make each of the bools above the opposite of each other
//!true is false
//!false is true
