package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {

	var result []string
	var s string

	for i := 1; i <= n; i++ {

		s = ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}

		if i%3 != 0 && i%5 != 0 {
			s = strconv.Itoa(i)
		}
		result = append(result, s)
	}

	return result
}

func main() {

	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}
