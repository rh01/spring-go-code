/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetric(root.Left, root.Right)
}

func isEual(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil && root2 != nil {
		return false
	} else if root1 != nil && root2 == nil {
		return false
	} 

	return root1.Val == root2.Val && isEual(root1.Left, root2.Right) && isEual(root1.Right, root2.Left)
}