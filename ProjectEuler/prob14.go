package main

func collatzSequence(n int) []int {
	sequence := make([]int, 0)

	for n > 1 {
		sequence = append(sequence, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	return append(sequence, 1)
}
