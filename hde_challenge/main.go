package main

//  2
//  4
//  3 -1 1 14
//  5
//  9 6 -53 32 16

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askForInteger() int64 {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	value, _ := strconv.ParseInt(text, 10, 64)
	return value
}

func getCases(N int64) int64 {

	if N > 1 {
		M := getCases(N - 1)
		println("Case ", M)
	}
	return askForInteger()
}

func main() {
	N := askForInteger()
	fmt.Println(N)

	getCases(N)
}
