package main

import (
	"fmt"
)

func main() {
	// pow := multiply("1099511627776", "1099511627776")
	// fmt.Println(pow, getDigitsSum(pow))
	pow := multiply("2", "1")
	n := 40
	for i := 2; i <= n; i++ {
		pow = multiply(pow, "2")

		fmt.Println("i=", i, pow)
	}

	fmt.Println("n=", n, getDigitsSum(pow), pow)
}
