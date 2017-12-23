package main

import (
	"fmt"
	"math"
)

// Brute force primality check
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	max := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// NOTE: Program is looking for non prime numbers between 109900 and 126900
//       in increments of 17
func main() {
	cnt := 0
	for i := 109900; i <= 126900; i += 17 {
		if isPrime(i) {
			fmt.Println(i)
		} else {
			cnt++
		}
	}
	fmt.Printf("DONE: %d\n", cnt)
}
