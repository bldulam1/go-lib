package main

import (
	"strconv"
)

func getNthDigit(number string, n int) int {
	if len(number) < int(n+1) {
		return 0
	}
	digit, _ := strconv.Atoi(number[n : n+1])
	return digit
}

func getLargeSum(numbers []string) string {
	carryOver := uint64(0)
	finalSum := ""
	for i := len(numbers[0]) - 1; i >= 0; i-- {
		tempSum := uint64(carryOver)
		for _, number := range numbers {
			digit := getNthDigit(number, int(i))
			tempSum += uint64(digit)
		}
		digit := tempSum % 10
		carryOver = (tempSum - digit) / 10

		finalSum = strconv.Itoa(int(digit)) + finalSum
	}
	if carryOver > 0 {
		finalSum = strconv.Itoa(int(carryOver)) + finalSum
	}

	return finalSum
}
