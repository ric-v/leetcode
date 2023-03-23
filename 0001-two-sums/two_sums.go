package twosums

// Two Sum returns the indices of the two numbers such that they add up to a specific target.
//
// Time complexity: O(n)
// Space complexity: O(n)
func TwoSum(nums []int, target int) []int {

	var m = make(map[int]int)

	// for each number in nums
	for i, num := range nums {

		// if the target minus the number is in the map
		// return the index of the number and the index of the target minus the number
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}

		// otherwise, add the number to the map
		m[num] = i
	}

	// if no two numbers add up to the target, return an empty array
	return []int{}
}
