package validparanthesis

// IsValid returns true if the string is a valid pair of brackets
// else returns false
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func IsValid(s string) bool {
	var stack []rune

	// iterate on each character in the string
	for _, char := range s {

		// if its an open bracket, push to stack and continue
		if isOpenBracket(char) {
			// push data to stack
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		// peek the last element in stack, if it's the matching pair of this closing bracket, then reduce stack by 1
		prevVal := stack[len(stack)-1]
		if prevVal == matchingPair(char) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	// if stack is not empty, string is not equal pair of brackets
	return len(stack) == 0
}

// returns if a character is an open bracket or not
func isOpenBracket(char rune) bool {
	return char == '(' || char == '{' || char == '['
}

// map of matching bracket pairs
func matchingPair(char rune) rune {
	switch char {
	case ')':
		return '('
	case '}':
		return '{'
	default:
		return '['
	}
}
