package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	return hasSubstructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func hasSubstructure(A, B *TreeNode) bool {
	if A == nil {
		if B == nil {
			return true
		}
		return false
	}

	if B == nil {
		return true
	}

	return A.Val == B.Val && hasSubstructure(A.Left, B.Left) && hasSubstructure(A.Right, B.Right)
}

func main() {

}
