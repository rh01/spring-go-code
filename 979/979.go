package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}

	L := dfs(node.Left, ans)  //负数表示左节点需要扣除的金币，需要从根节点搬金币下去；正数表示多出的金币，需要将多出的金币搬上去；
	R := dfs(node.Right, ans) //负数表示左节点需要扣除的金币，需要从根节点搬金币下去；正数表示多出的金币，需要将多出的金币搬上去；

	*ans = *ans + abs(L) + abs(R) // //计算路径（金币数的绝对值就是路径）

	return node.Val + L + R - 1 //返回节点的剩余的金币量（已经扣除本身以及左、右子节点需要的）
}

func abs(i0 int) int {
	if i0 < 0 {
		return -i0
	}
	return i0
}

func main() {

}
