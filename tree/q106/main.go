package main

/**
 * Definition for a binary tree node.
 * */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}

	v := postorder[len(postorder)-1]
	idx := findIndex(inorder, v)
	root := &TreeNode{
		v,
		buildTree(inorder[:idx], postorder[:idx]),
		buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
	return root
}

func findIndex(arr []int, targetNum int) int {
	for i, j := range arr {
		if j == targetNum {
			return i
		}
	}
	return -1
}

func main() {
	
}