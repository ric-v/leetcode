package main

import "fmt"

func dominantIndex(nums []int) int {

	var maxIndex int
	for i, n := range nums {
		if nums[maxIndex] < n {
			maxIndex = i
		}
	}

	for i, n := range nums {
		if i != maxIndex && !(nums[maxIndex] >= 2*n) {
			return -1
		}
	}

	return maxIndex
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
}
