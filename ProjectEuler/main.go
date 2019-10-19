package main

import (
	"bufio"
	"log"
	"os"
)

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

func main() {

	// product := specialPythagoreanTriple(1000)

	contents := readTextFile("./files/series.txt")
	largestProduct(contents, 13)
}
