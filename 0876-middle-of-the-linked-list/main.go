package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) len() int {

	var tmpHead = l
	var c int

	for tmpHead.Next != nil {
		tmpHead = tmpHead.Next
		c++
	}
	c++
	return c
}

func middleNode(head *ListNode) *ListNode {

	var tmpHead = head
	var length = head.len()

	var c int
	for tmpHead.Next != nil && c < length/2 {
		tmpHead = tmpHead.Next
		c++
	}
	return tmpHead
}

func main() {
	fmt.Println(
		middleNode(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		),
	)
	fmt.Println(
		middleNode(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		),
	)
}
