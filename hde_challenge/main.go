package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func askForInteger() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	value, _ := strconv.ParseInt(text, 10, 64)
	return int(value)
}

func askForArray() []string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	return strings.Split(text, " ")
}

func getSquaresSum(array []string, arrayLen int) int {
	sum := 0

	for index := 0; index < int(arrayLen); index++ {
		text := array[index]
		val, _ := strconv.ParseInt(text, 10, 64)
		if val > 0 {
			sum += int(val * val)
		}
	}

	return sum
}

func getCases(N int, sums []int) {
	if N == 0 {
		return
	}
	arrayLen := askForInteger()
	array := askForArray()
	sum := getSquaresSum(array[:], arrayLen)
	sums[len(sums)-N] = sum
	getCases(N-1, sums)

}

func main() {
	N := askForInteger()
	sums := make([]int, N)
	getCases(N, sums)

	for i := 0; i < N; i++ {
		println(sums[i])
	}
}
