package main

func main() {
	str := readTextFile("./files/prob13.txt")
	numLen := 50
	numbers := make([]string, 0)

	for i := 0; i < len(str); i += numLen {
		numbers = append(numbers, str[i:i+numLen])
	}

	sum := getLargeSum(numbers)
	println(sum)
}
