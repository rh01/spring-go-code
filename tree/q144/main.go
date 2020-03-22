package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res = make([]int, 0)
	walk(root, res)
	return res
}

func walk(node *TreeNode, res []int) {
	if node == nil {
		return
	}

	res = append(res, node.Val)
	walk(node.Left, res)
	walk(node.Right, res)
	return 
} 

func main() {

}
