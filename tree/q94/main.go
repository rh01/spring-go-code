package main

// TreeNode node of tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack = make([]*TreeNode, 0)
	var res = make([]int, 0)
	node := root
	for len(stack) != 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			// pop
			node =stack[len(stack)-1]
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]

			node = node.Right
		}
	}
	return res
}

func main() {

}
