/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func check(node *TreeNode, min *int, max *int) bool {
	if nil == node {
		return true
	}

	if (nil != min && node.Val <= *min) || (nil != max && node.Val >= *max) {
		return false
	}

	return check(node.Right, min, &node.Val) && check(node.Left, &node.Val, max)
}

func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}

	return check(root.Left, nil, &root.Val) && check(root.Right, nil, &root.Val)
}