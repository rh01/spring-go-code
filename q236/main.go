package main

/**
 * Definition for TreeNode.
 * */
type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var m = make(map[*TreeNode]*TreeNode)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	m[root] = nil
	for {
		_, ok1 := m[p]
		_, ok2 := m[q]
		if ok1 && ok2 {
			break
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
			m[node.Left] = node
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			m[node.Right] = node
		}

	}

	var path = make(map[*TreeNode]bool)
	for p != nil {
		path[p] = true
		p = m[p]
	}

	for q != nil {
		if path[q] {
			break
		}
		q = m[q]
	}

	return q

}

func main() {

}
