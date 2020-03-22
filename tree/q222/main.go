package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var hl, hr uint

	l, r := root, root

	for l != nil {
		hl++
		l = l.Left
	}

	for r != nil {
		hr++
		r = r.Right
	}

	if hl == hr {
		return 1<<hl - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
func main() {

}
