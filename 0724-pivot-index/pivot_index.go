package pivotindex

// PivotIndex returns the middle index of the array where the sum of the left and right side of the index is equal.
//
// Time complexity: O(n)
// Space complexity: O(1)
func PivotIndex(num []int) int {
	// We will use two variables to keep track of the left and right sum
	var left = 0
	var right = sum(num[1:])
	// Loop through the array and compare the left and right sum
	for i, n := range num {
		// If the left and right sum is equal, return the index
		if left == right {
			return i
		}
		// Add the current index value to the left sum
		left += n
		// If the array length is greater than the current index + 1, subtract the next index from the right sum
		if len(num) > i+1 {
			right -= num[i+1]
		}
	}
	// If no index is found, return -1
	return -1
}

// sum() calculates the sum of a slice of integers
// using a range loop.
func sum(num []int) int {
	// Declare a variable to hold the sum
	var sum int
	// Loop through each element in the slice
	for _, n := range num {
		// Add the current element to the sum
		sum += n
	}
	// Return the sum
	return sum
}
