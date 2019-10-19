package main

func specialPythagoreanTriple(sum int) int {
	for a := 1; a < sum; a++ {
		for b := 1; b < sum; b++ {
			c := sum - a - b
			if c*c == a*a+b*b {
				return a * b * c
			}
		}
	}
	return 0
}
