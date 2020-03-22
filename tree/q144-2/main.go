package main

/**
 * Definition for a binary tree node.
 *  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// 利用栈来解决这个问题
	var stack = make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}

	var node = new(TreeNode)
	var res = make([]int, 0)
	for stack != nil || len(stack) == 0 {
		node = stack[0]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func main() {

}
