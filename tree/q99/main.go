package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var f, s, p, n, prev *TreeNode = nil, nil, nil, root, nil
	for n != nil {
		if n.Left == nil {
			if p != nil && p.Val > n.Val {
				if f == nil {
					f = p
				}
				s = n // f, s
			}
		} else {
			prev = n.Left
			for prev.Right != nil && prev.Right != n {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right, n = n, n.Left
			} else {
				prev.Right = nil
				if p.Val > n.Val {
					if f == nil {
						f = p
					}
					s = n
				}
				p, n = n, n.Right
			}
		}
	}
	f.Val, s.Val = s.Val, f.Val
}

func main() {

}
