package lcof

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	var preorder = []int{3, 9, 20, 15, 7}
	var inorder = []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)

	res := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	if isTwoTreeEqual(res, root) {
		t.Errorf("want %v, but got %v\n", res, root)
	}
}

// 判断两棵树是否相等
func isTwoTreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	var isLeft, isRight bool
	if t1.Left != nil && t2.Left != nil {
		isLeft = isTwoTreeEqual(t1.Left, t2.Left)
	}

	if t1.Right != nil && t2.Right != nil {
		isRight = isTwoTreeEqual(t1.Right, t2.Right)
	}

	return t1.Val == t2.Val && isLeft && isRight
}

func TestBinarysearch(t *testing.T) {
	prepareData := []int{9, 3, 15, 20, 7}
	target := 3
	if res := searchIndex(prepareData, target); res != 0 {
		t.Errorf("want %v, but got %v\n", 1, res)
	}
}
