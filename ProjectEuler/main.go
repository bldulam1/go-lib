package main

import "fmt"

func main() {
	divisorsLen := 0
	index := 12370
	maxDivisorsLen := 500
	for divisorsLen <= maxDivisorsLen {
		an := arithmeticSum(index, 1, 1)
		divisors := getDivisors(an)
		divisorsLen = len(divisors)
		fmt.Println(index, an, divisorsLen)
		index++
	}
}
