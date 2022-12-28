package main

import "fmt"

func addOne(digits []int, index int) []int {
	if digits[index]+1 >= 10 {
		digits[index] = 0

		if index-1 < 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
		fmt.Println(digits)
		return addOne(digits, index-1)
	}
	digits[index] += 1
	return digits
}

func plusOne(digits []int) []int {
	return addOne(digits, len(digits)-1)
}

func main() {

	// fmt.Println(plusOne([]int{1, 2, 3}))
	// fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
}
