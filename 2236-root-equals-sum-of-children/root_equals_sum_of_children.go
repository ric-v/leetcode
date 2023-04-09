package rootequalssumofchildren

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rootEqualsSumOfChildren returns true if the root node equals the sum of its children.
//
// Time complexity: O(1)
// Space complexity: O(1)
func RootEqualSumOfChildren(root *TreeNode) bool {
	return (root.Left.Val + root.Right.Val) == root.Val
}
