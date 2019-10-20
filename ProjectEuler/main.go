package main

import "fmt"

func main() {

	maxCollatzLen := 0
	maxCollatzNum := 0

	for i := 13; i < 1e6; i++ {
		collatz := collatzSequence(i)
		lcLen := len(collatz)

		if lcLen > maxCollatzLen {
			maxCollatzLen = lcLen
			maxCollatzNum = i
		}
	}

	fmt.Println(maxCollatzNum, maxCollatzLen)
}
