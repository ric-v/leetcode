package main

import "fmt"

func sum(num []int) int {
	var sum int
	for _, n := range num {
		sum += n
	}
	return sum
}

func findMiddleIndex(num []int) int {

	var left = 0
	var right = sum(num[1:])
	for i, n := range num {

		if left == right {
			return i
		}
		left += n
		if len(num) > i+1 {
			right -= num[i+1]
		}
	}
	return -1
}

func main() {

	fmt.Println(findMiddleIndex([]int{1, 7, 3, 6, 5, 6}))

	fmt.Println(findMiddleIndex([]int{1, 2, 3}))

	fmt.Println(findMiddleIndex([]int{2, 1, -1}))
}
