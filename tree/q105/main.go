package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	val := preorder[0]
	idx := getIndex(inorder, val)
	root := &TreeNode{
		Val:   val,
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
	return root
}

func getIndex(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {

}
