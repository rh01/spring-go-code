package lcof

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//  这个是迭代，需要反复练习
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 创建root节点
	root := &TreeNode{preorder[0], nil, nil}
	// 创建辅助栈
	stack := make([]*TreeNode, 0)
	// 将根节点入栈
	stack = append(stack, root)
	// 初始化中序索引
	var inorderIdx int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIdx] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIdx] {
				node = stack[len(stack)-1]
				// pop
				stack = stack[:len(stack)-1]
				inorderIdx++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

// 递归的建树过程
func buildTreev1(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 || inorder == nil || len(inorder) == 0 {
		return nil
	}

	rVal := preorder[0]
	idx := searchIndex(inorder, rVal)
	if idx == -1 {
		return nil
	}
	root := &TreeNode{
		Val:   rVal,
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}

	return root
}

// 下面是一个二分查找过程
func searchIndex(inorder []int, target int) int {
	// 先做一个异常处理
	if inorder == nil || len(inorder) == 0 {
		return -1
	}
	for idx, v := range inorder {
		if v == target {
			return idx
		}
	}
	return -1
}
