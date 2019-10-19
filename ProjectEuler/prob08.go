package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getProduct(str string) uint64 {
	prod := uint64(1)
	count := len(str)
	for i := 0; i < count; i++ {
		num, _ := strconv.Atoi(string(str[i]))
		if num > 0 {
			prod *= uint64(num)
		} else {
			prod = 0
			break
		}
	}
	return prod
}

func readTextFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	contents := ""

	for scanner.Scan() {
		contents += scanner.Text()
	}
	return contents
}

func largestProduct(str string, num int) {
	stringLen := len(str)
	maxProduct := uint64(0)
	for i := 0; i < stringLen-num; i++ {
		subset := str[i : i+num]
		prod := getProduct(subset)
		if prod > maxProduct {
			maxProduct = prod
		}
	}

	print(maxProduct)
}
