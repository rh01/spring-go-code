package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil && root.Right != nil {
		root.Left.Val, root.Right.Val = root.Right.Val, root.Left.Val
	} else if  root.Left != nil && root.Right == nil {
		root.Left, root.Right = nil, root.Left
	} else if  root.Left == nil && root.Right != nil {
		root.Right, root.Left = nil, root.Right
	} else {
		return root
	}

	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)
	return root
}
