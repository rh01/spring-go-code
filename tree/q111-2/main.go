package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		var tmpNodes = make([]*TreeNode, 0)
		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			} else if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			} else if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}
		}
		queue = append(queue, tmpNodes...)
		depth++
	}
	return depth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
