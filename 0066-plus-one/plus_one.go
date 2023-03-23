package plusone

// PlusOne returns the digits array with 1 added to it
//
// Time complexity: O(n)
// Space complexity: O(1)
func PlusOne(digits []int) []int {

	var carryForward bool

	// iterate over each digit in the array
	for i := len(digits) - 1; i >= 0; i-- {

		// add 1 to the digit
		n := digits[i] + 1

		// if n > 9, then set n to 0 and set carryForward to true
		if n > 9 {
			n = 0
			carryForward = true
		}

		// set the digit to n
		digits[i] = n

		// if carryForward is false, then break
		if n > 0 {
			carryForward = false
			break
		}
	}

	// if carryForward is true, then add 1 to the beginning of the array
	if carryForward {
		digits = append([]int{1}, digits...)
	}
	return digits
}
