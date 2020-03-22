package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	var res = make([]int, 0)
	inOrderFound(root, &res)
	return res[k-1]
}

func inOrderFound(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrderFound(root.Left, res)
	*res = append(*res, root.Val)
	inOrderFound(root.Right, res)
}

func main() {

}
