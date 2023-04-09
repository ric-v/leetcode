package runningsumof1darray

// PrefixSum returns the running sum of nums.
// 
// Time complexity: O(n)
// Space complexity: O(1)
func PrefixSum(nums []int) []int {
	// loop through the array
	for i := 1; i < len(nums); i++ {
		// For each element, add the previous element to itself
		nums[i] += nums[i-1]
	}
	// Return the array
	return nums
}
