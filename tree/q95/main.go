package tree

/**
 * Definition for a binary tree node.
 * */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GenerateTrees dcsfffff
func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	return calc(1, n)
}

func calc(start, end int) []*TreeNode {
	var res = make([]*TreeNode, 0)

	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ {
		ls, rs := calc(start, i-1), calc(i+1, end)

		for _, l := range ls {
			for _, r := range rs {
				res = append(res, &TreeNode{i, l, r})
			}
		}
	}

	return res
}

func main() {

}
