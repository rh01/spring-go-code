/**
 * Definition for a binary tree node.
 * */ 
 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func diameterOfBinaryTree(root *TreeNode) int {
	return depth(root.Left) + depth(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(depth(node.Left), depth(node.Right))+1
}

func max(i0, i1 int ) int {
	if i0 > i1 {
		return i0
	}
	return i1
}