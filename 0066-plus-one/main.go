package main

import "fmt"

func plusOne(digits []int) []int {

	var carryForward bool
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + 1

		if n > 9 {
			n = 0
			carryForward = true
		}
		digits[i] = n

		if n > 0 {
			carryForward = false
			break
		}
	}

	if carryForward {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {

	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{8, 9, 9, 9}))
}
