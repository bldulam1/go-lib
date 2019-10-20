package main

func arithmeticSum(n int, a1 int, increment int) int {
	return (2*a1 + (n-1)*increment) * n / 2
}

func getDivisors(n int) []int {
	divisors := make(map[int]int)
	divisors[1] = n

	factors := make([]int, 0)

	if n > 1 {
		factors = append(factors, 1)
		factors = append(factors, n)
		for i := 2; i < n/2; i++ {
			if n%i == 0 && divisors[n/i] == 0 {
				divisors[i] = n / i
				factors = append(factors, i)
				if divisors[i] != i {
					factors = append(factors, divisors[i])
				}
			}
		}
	}

	return factors
}
