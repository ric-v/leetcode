package addtwonums

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers adds two numbers represented by linked lists
// and returns the sum as a linked list.
//
// Time complexity: O(max(m, n))
// Space complexity: O(max(m, n))
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		carry int
		head  *ListNode
		tail  *ListNode
	)

	// while l1 or l2 or carry is not zero
	// add the values of l1 and l2 and carry
	for l1 != nil || l2 != nil || carry != 0 {

		var (
			v1 int
			v2 int
		)

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		// if the sum is greater than 10, set carry to 1 and set the sum to the remainder
		// otherwise, set carry to 0
		// create a new node with the sum
		sum := v1 + v2 + carry
		carry = sum / 10

		node := &ListNode{Val: sum % 10}

		// if head is nil, set head and tail to the new node
		// otherwise, set tail.Next to the new node and set tail to the new node
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}
