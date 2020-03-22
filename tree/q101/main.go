package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isEqual(root.Left, root.Right)
}

func isEqual(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil && node2 != nil {
		return false
	}

	if node2 == nil  && node1 != nil {
		return false
	}

	return node1.Val == node2.Val && isEqual(node1.Left, node2.Right) && isEqual(node1.Right, node2.Left)
}

func main() {

}
