package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var level int
	var res = make([][]int, 0)
	if root == nil {
		return res
	}

	var stack = make([]*TreeNode, 0)
	stack = append(stack, root)
	var node *TreeNode
	for len(stack) != 0 {
		var tmpNodes = make([]*TreeNode, 0)
		var tmpRes = make([]int, 0)
		for len(stack) != 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if level%2 == 0 {
				// node = stack[0]
				// stack = stack[1:]

				if node.Left != nil {
					tmpNodes = append(tmpNodes, node.Left)
				}

				if node.Right != nil {
					tmpNodes = append(tmpNodes, node.Right)
				}
			} else {
				//  node = stack[0]
				// stack = stack[1:]

				if node.Right != nil {
					tmpNodes = append(tmpNodes, node.Right)
				}

				if node.Left != nil {
					tmpNodes = append(tmpNodes, node.Left)
				}
			}
			tmpRes = append(tmpRes, node.Val)
		}
		level++
		stack = append(stack, tmpNodes...)
		res = append(res, tmpRes)
	}
	return res
}
func main() {

}
