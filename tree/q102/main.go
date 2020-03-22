package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue = make([]*TreeNode, 0)
	var res = make([][]int, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		// nList := queue[0]
		var tmp = make([]int, 0)
		var tmpNodes = make([]*TreeNode, 0)
		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			}

			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}

			tmp = append(tmp, node.Val)
		}
		queue = append(queue, tmpNodes...)
		res = append(res, tmp)
	}
	return res
}

func main() {

}
