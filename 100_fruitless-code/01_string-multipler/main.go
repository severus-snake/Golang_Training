package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	firstLetter := strconv.FormatInt(110, 2)
	secondLetter := strconv.FormatInt(111, 2)
	bytecode := firstLetter + secondLetter

	result := float64(0)
	columnValue := float64(1)

	for index := len(bytecode)-1; index >= 0; index-- {
		if (bytecode[index] != '0'){
			result = result + columnValue
		}
		columnValue *= 2
	}

	result = math.Pow(result, 2)

	//10	2	2	0 result

	//1	0	1	0 bytecode[index]
	//8	4	2	1 columnValue

	fmt.Println(firstLetter)
	fmt.Println(secondLetter)
	fmt.Println(bytecode)
	fmt.Println(result)
}

