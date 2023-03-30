package rangesumqueryimmutable

// NumArray is a struct that contains the cumulative sum of the elements of the nums array (Prefix Sum).
type NumArray struct {
	Arr []int
}

// Constructor initializes the object with the integer array nums.
// This calculates the cumulative sum of the elements of the nums array (Prefix Sum).
//
// Time complexity: O(n)
// Space complexity: O(n)
func Constructor(nums []int) NumArray {
	var arr []int = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		arr = append(arr, nums[i]+arr[i-1])
	}
	return NumArray{arr}
}

// SumRange returns the sum of the elements of the nums array in the range [left, right] inclusive.
//
// Time complexity: O(1)
// Space complexity: O(1)
func (numArr *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return numArr.Arr[right]
	} else {
		return numArr.Arr[right] - numArr.Arr[left-1]
	}
}
