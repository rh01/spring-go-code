package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walkminDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return walk(root, 1)

}

// var level int
func walk(root *TreeNode, level int) int {
	if root.Left == nil && root.Right == nil {
		return level
	} else if root.Left == nil && root.Right != nil {
		return walk(root.Right, level+1)
	} else if root.Left != nil && root.Right == nil {
		return walk(root.Left, level+1)
	} else {
		leftDepth := walk(root.Left, level+1)
		rightDepth := walk(root.Right, level+1)
		return min(leftDepth, rightDepth)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
