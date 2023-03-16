package main

import (
	"fmt"
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
func longestCommonPrefix(strs []string) (prefix string) {
	if len(strs) < 1 || len(strs) > 0 && len(strs[0]) < 1 {
		return
	}

	var l int = 1
	prefix = strs[0][:l]
	for i := 0; i < len(prefix); i++ {

		for j, str := range strs {

			if j == 0 {
				continue
			}

			if !strings.HasPrefix(str, prefix) {
				if len(prefix) > 1 && i != 0 {
					return prefix[:i]
				}
				return ""
			}
		}
		l++

		if l <= len(strs[0]) {
			prefix = strs[0][:l]
		} else {
			return
		}
	}
	return
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}
