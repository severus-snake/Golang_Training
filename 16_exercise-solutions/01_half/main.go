package main

import "fmt"

func sum_prime(slice []int) int {
	var sum int
	for _, v := range slice {
		if v >= sum{
			sum += v
		}
	}
	return sum
}

func prime_numbers(n int) []int {
	var slice []int
	var result int

	for n > 1 {
		result = find_factors(n)
		slice = append(slice, result)
		n = n/find_factors(n)
	}
	return slice
}

func find_factors(n int) int {
	x := 2
	for i:= 2; n % i !=0; i++ {
		x = i + 1
	}
	return x
}

func main() {
	var n int
	fmt.Print("Please enter a number:")
	fmt.Scan(&n)
	prime_slice := prime_numbers(n)
	sum := sum_prime(prime_slice)
	fmt.Println("Prime Numbers for your number:", prime_slice)
	fmt.Println("The sum of those prime numbers:", sum)

}


/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/