package fizzbuzz

import (
	"strconv"
)

// FizzBuzz returns a slice of strings from 1 to n
// where multiples of 3 are replaced by "Fizz"
// and multiples of 5 are replaced by "Buzz"
// and multiples of 3 and 5 are replaced by "FizzBuzz"
//
// Time complexity: O(n)
// Space complexity: O(n)
func FizzBuzz(n int) []string {

	var result []string
	var s string

	// iterate over each number from 1 to n
	for i := 1; i <= n; i++ {

		// if i is a multiple of 3, append "Fizz" to s
		s = ""
		if i%3 == 0 {
			s += "Fizz"
		}

		// if i is a multiple of 5, append "Buzz" to s
		if i%5 == 0 {
			s += "Buzz"
		}

		// if i is not a multiple of 3 or 5, append the number to s
		if i%3 != 0 && i%5 != 0 {
			s = strconv.Itoa(i)
		}
		result = append(result, s)
	}
	return result
}
