package main

import "math"

func primeFactors(n int) map[int]int {
	factors := make(map[int]int)

	i := 2
	if n%i == 0 {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	for i = 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			for n%i == 0 {
				factors[i]++
				n /= i
			}
		}
	}

	if n > 2 {
		factors[n]++
	}
	return factors
}

func smallestMultiple(min int, max int) float64 {
	pFactors := make(map[int]int)
	multiple := float64(1)

	for num := min; num <= max; num++ {
		factors := primeFactors(num)

		for key, value := range factors {
			if pFactors[key] < value {
				pFactors[key] = value
			}
		}
	}

	for key, value := range pFactors {
		multiple *= math.Pow(float64(key), float64(value))
	}

	return multiple
}
