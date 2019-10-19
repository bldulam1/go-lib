package main

func sumSquareDifference(n int) int {
	A := n * (n + 1) / 2
	B := (2*n + 1) / 3
	return A * (A - B)
}
