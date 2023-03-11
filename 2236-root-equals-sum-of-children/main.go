package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return (root.Left.Val + root.Right.Val) == root.Val
}

func main() {
	fmt.Println(checkTree(&TreeNode{Val: 10, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}))
	fmt.Println(checkTree(&TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}))
}
