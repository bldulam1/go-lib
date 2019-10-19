package main

func summationOfPrimes(n int) uint64 {
	if n < 2 {
		return 1
	} else if n < 3 {
		return 2
	}

	num := 3
	sum := uint64(2)

	for num <= n {
		if isPrime(num) {
			sum += uint64(num)
			// println(sum)
		}
		num += 2
	}

	return sum
}
