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
	var stack = make([]*TreeNode, 0)
	var res = make([]int, 0)
	if root == nil {
		return res
	}

	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverse(&res)
	return res
}

func reverse(arr *[]int) {
	start, end := 0, len(*arr)-1
	for start < end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start++
		end--
	}
}

func main() {

}
