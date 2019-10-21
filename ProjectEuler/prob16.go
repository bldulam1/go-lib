package main

import (
	"strconv"
	"strings"
)

func getDigits(num string) []uint64 {
	LEN := 8

	numUint := make([]uint64, 0)
	numLen := len(num)

	if numLen%LEN > 0 {
		num = strings.Repeat("0", LEN-numLen%LEN) + num
	}

	for i := 0; i < numLen; i += LEN {
		digits, _ := strconv.ParseUint(num[i:i+LEN], 10, 64)
		numUint = append(numUint, digits)
	}

	return numUint
}

func multiply(a string, b string) string {
	LEN := 9

	aDigits := getDigits(a)
	bDigits := getDigits(b)

	// fmt.Println(aDigits, bDigits)

	addends := make([]string, 0)
	addendLen := 9 * (len(aDigits) + len(bDigits))

	for iA := len(aDigits) - 1; iA >= 0; iA-- {
		aDigit := aDigits[iA]
		addend := ""
		carryOver := uint64(0)
		leftOver := uint64(0)

		for iB := len(bDigits) - 1; iB >= 0; iB-- {
			bDigit := bDigits[iB]
			prod := aDigit*bDigit + uint64(carryOver)
			leftOver = uint64(prod % 1e9)
			carryOver = uint64(prod / 1e9)
			strLeftOver := strconv.Itoa(int(leftOver))
			strLeftOver = strings.Repeat("0", LEN-len(strLeftOver)) + strLeftOver
			addend = strLeftOver + addend
		}

		if carryOver > 0 {
			strCarryOver := strconv.Itoa(int(carryOver))
			strCarryOver = strings.Repeat("0", LEN-len(strCarryOver)) + strCarryOver
			addend = strCarryOver + addend
		}

		if len(aDigits)-iA > 1 {
			addend += strings.Repeat("0", LEN*(len(aDigits)-iA-1))
		}

		addend = strings.Repeat("0", addendLen-len(addend)) + addend
		addends = append(addends, addend)
	}

	largeSum := ""
	if len(addends) < 2 {
		largeSum = addends[0]
	} else {
		largeSum = getLargeSum(addends)
	}

	lsDigit, _ := strconv.Atoi(largeSum[0:1])
	for lsDigit < 1 {
		largeSum = largeSum[1:]
		lsDigit, _ = strconv.Atoi(largeSum[0:1])
	}

	return largeSum
}

func getDigitsSum(num string) int {
	sum := 0
	for i := range num {
		sum += getNthDigit(num, i)
	}
	return sum
}
