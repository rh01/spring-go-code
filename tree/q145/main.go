package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	var res = make([]int, 0)

	walk(root, &res)
	return res
}

func walk(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	walk(node.Left, res)
	walk(node.Right, res)
	*res = append(*res, node.Val)
	return
}

func main() {

}
