package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var m = make(map[int]int)
	for i, n1 := range nums {

		if n2, ok := m[target-n1]; ok {
			return []int{n2, i}
		}
		m[n1] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
