/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return [][]int{}
	}

	q := []*TreeNode{}
	res := [][]int{}
	for 0 != len(q) {
		pq := []*TreeNode{}
		qRes := []int{}
		for 0 != len(q) {
			node := q[0]
			q = q[1:]
			qRes = append(qRes, node.Val)
			if node.Left != nil {
				pq = append(pq, node.Left)
			}
			if node.Right != nil {
				pq = append(pq, node.Right)
			}
		}
		copy(q,  pq)
		res = append(res, qRes)
	}
	return res
}