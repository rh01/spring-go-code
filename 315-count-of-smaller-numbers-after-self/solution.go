package main

import "fmt"

const (
	MAX_LEN = 1000
)

func buildTree(arr []int, tree *[MAX_LEN]int, node, start, end int) {
	// 递归的出口
	if start == end {
		tree[node] = arr[start]
	} else {
		mid := start + (end-start)/2
		left_node := 2*node + 1
		right_node := 2*node + 2

		buildTree(arr, tree, left_node, start, mid)
		buildTree(arr, tree, right_node, mid+1, end)

		tree[node] = tree[left_node] + tree[right_node]
	}
}

// updateTree 将idx的数据改成val
func updateTree(arr []int, tree *[MAX_LEN]int, node, start, end int, idx, value int) {
	// 边界条件
	if start == end {
		arr[idx] = value
		tree[node] = value
	} else {
		mid := start + (end-start)/2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		if idx >= start && idx <= mid {
			updateTree(arr, tree, leftNode, start, mid, idx, value)
		} else {
			updateTree(arr, tree, rightNode, mid+1, end, idx, value)
		}
		// base case
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}

func queryTree(arr []int, tree *[MAX_LEN]int, node, start, end int, L, R int) int {
	fmt.Printf("L=[%d], R=[%d]\n",start, end)
	// 这个是边界条件
	if L > end || R < start {
		return 0
	} else if start >= L && R >= end {
		return tree[node]
	} else if start == end {
		return tree[node]
	}

	// 还是计算中间节点，然后比较
	mid := start + (end-start)/2
	leftNode := 2*node + 1
	rightNode := 2*node + 2

	leftSum := queryTree(arr, tree, leftNode, start, mid, L, R)
	rightSum := queryTree(arr, tree, rightNode, mid+1, end, L, R)
	return leftSum + rightSum
}

func main() {

	array := []int{1, 3, 5, 7, 9, 11}
	size := 6
	tree := [MAX_LEN]int{0}
	buildTree(array, &tree, 0, 0, size-1)
	for i := 0; i < 15; i++ {
		fmt.Printf("tree[%d] = %d\n", i, tree[i])
	}

	fmt.Println()

	updateTree(array, &tree, 0, 0, size-1, 4, 6)
	for i := 0; i < 15; i++ {
		fmt.Printf("tree[%d] = %d\n", i, tree[i])
	}

	fmt.Println()
	fmt.Println(queryTree(array, &tree, 0, 0, size-1, 2, 5))
}
