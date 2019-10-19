package main

func isPrime(n int) bool {

	for i := 2; i < n/2; i++ {
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
