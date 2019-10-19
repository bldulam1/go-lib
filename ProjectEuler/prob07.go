package main

import "math"

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func nthPrime(n int) int {
	i := 2
	num := 3

	for i <= n {
		if isPrime(num) {
			i++
		}
		num += 2
	}
	return num - 2
}
