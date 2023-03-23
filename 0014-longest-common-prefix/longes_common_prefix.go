package longestcommonprefix

import (
	"strings"
)

/*
Input: strs = ["flower", "flow", "flight"]

Step 1: Check input
  - len(strs) = 3
  - len(strs[0]) = 6
  - Continue to next step

Step 2: Initialize variables
  - prefix = "f"

Step 3: Iterate over each character of prefix
  - i = 0:
  - Iterate over each string in strs
  - j = 0: continue
  - j = 1:
  - strings.HasPrefix("flow", "f") = true
  - j = 2:
  - strings.HasPrefix("flight", "f") = true
  - l = 2
  - prefix = "fl"
  - i = 1: since prefix now increased from 'f' to 'fl', i can be 1
  - Iterate over each string in strs
  - j = 0: continue
  - j = 1:
  - strings.HasPrefix("flow", "fl") = true
  - j = 2:
  - strings.HasPrefix("flight", "fl") = true
  - l = 3
  - prefix = "flo"
  - i = 2: since prefix now increased from 'fl' to 'flo', i can be 2
  - Iterate over each string in strs
  - j = 0: continue
  - j = 1:
  - strings.HasPrefix("flow", "flo") = true
  - j = 2:
  - strings.HasPrefix("flight", "flo") = false
  - Return "flo"

Output: "flo"
*/
// LongestCommonPrefix returns the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func LongestCommonPrefix(strs []string) (prefix string) {

	// Check input and return if input is invalid or empty
	if len(strs) < 1 || len(strs) > 0 && len(strs[0]) < 1 {
		return
	}

	var l int = 1
	// Initialize prefix with first character of first string
	prefix = strs[0][:l]

	// Iterate over each character of prefix
	for i := 0; i < len(prefix); i++ {

		// Iterate over each string in strs
		for j, str := range strs {

			// Skip first string as it is already used to initialize prefix
			if j == 0 {
				continue
			}

			// Check if prefix is a prefix of current string
			if !strings.HasPrefix(str, prefix) {

				// If prefix is not a prefix of current string, return prefix
				if len(prefix) > 1 && i != 0 {
					return prefix[:i]
				}
				return ""
			}
		}
		// If prefix is a prefix of all strings, increase length of prefix by 1
		l++

		// If length of prefix is less than or equal to length of first string, update prefix
		if l <= len(strs[0]) {
			prefix = strs[0][:l]
		} else {
			return
		}
	}
	return
}
