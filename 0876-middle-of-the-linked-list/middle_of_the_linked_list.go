package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode returns the middle node of a linked list.
//
// Time complexity: O(n)
// Space complexity: O(1)
func MiddleNode(head *ListNode) *ListNode {

	// get length of linkedlist
	var length = head.len()

	// create counter to count each step
	var c int

	// iterate through linkedlist
	for head.Next != nil && c < length/2 {
		head = head.Next
		c++
	}
	return head
}

func (l *ListNode) len() int {

	// Create a temporary head to iterate through the list
	var tmpHead = l

	// Create a counter to keep track of how many nodes are in the list
	var c int

	// Iterate through the list until we reach the last node
	for tmpHead.Next != nil {
		tmpHead = tmpHead.Next
		c++
	}

	// Add 1 to the counter and return it
	c++
	return c
}
