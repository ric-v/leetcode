package richestcustomerwealth

// MaximumWealth returns the maximum wealth of the customers.
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func MaximumWealth(accounts [][]int) int {
	// Initialize max to 0.
	var max int

	// Iterate over each account.
	for _, a := range accounts {
		// Sum the account's transactions.
		total := sum(a)

		// Update max if total is larger.
		if total > max {
			max = total
		}
	}
	return max
}

// This function sums all numbers in a slice of integers
func sum(act []int) (total int) {
	// iterate over integers in the slice
	for _, n := range act {
		// add each integer to total
		total += n
	}
	// return the sum of all integers in the slice
	return total
}
