package main

/**
 * Definition for a binary tree node.
 * */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(getDepth(root.Left)-getDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(getDepth(node.Left), getDepth(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
