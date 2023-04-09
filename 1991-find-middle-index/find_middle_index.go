package findmiddleindex

// MiddleIndex returns the middle index of the array where the sum of the left and right side of the index is equal.
//
// Time complexity: O(n)
// Space complexity: O(1)
func MiddleIndex(num []int) int {

	var left = 0             // left is the sum of all the numbers before index i
	var right = sum(num[1:]) // right is the sum of all the numbers after index i
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

// sum() calculates the sum of a slice of integers
func sum(num []int) int {
	var sum int
	for _, n := range num {
		sum += n
	}
	return sum
}
