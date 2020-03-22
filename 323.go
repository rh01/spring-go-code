/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
	q = append(q, root)
	level := 1
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
		q = append(q, pq...)
		if level % 2 != 0 {
			res = append(res, qRes)
		} else {
            reverse(qRes)
		}
		level++
	}
	return res
}

func reverse(qRes []int) {
	l, r := 0, len(qRes) - 1
	for l <= r {
		qRes[l], qRes[r] = qRes[r], qRes[l]
		l++
		r--
	}
}