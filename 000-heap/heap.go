package main

import "fmt"

// heapify 表示调整堆的过程，n表示arr的大小，从上往下做
func heapify(arr []int, n int, i int) {
	if i >= n {
		return
	}

	c1 := 2*i + 1
	c2 := 2*i + 2

	maxNode := i

	if c1 < n && arr[c1] > arr[maxNode] {
		maxNode = c1
	}

	if c2 < n && arr[c2] > arr[maxNode] {
		maxNode = c2
	}

	if maxNode != i {
		arr[i], arr[maxNode] = arr[maxNode], arr[i]
		heapify(arr, n, maxNode)
	}
}

func buildHeap(tree []int, n int) {
	lastNode := n - 1
	parentNode := (lastNode - 1) / 2

	// var i int
	for i := parentNode; i >= 0; i-- {
		heapify(tree, n, i)
	}

}

func heapsort(arr *[]int, n int) {
	buildHeap(*arr, n)
	for i := n - 1; i != 0; i-- {
		(*arr)[0], (*arr)[i] = (*arr)[i], (*arr)[0]
		heapify(*arr, i, 0)
	}
}

func main() {
	tree := []int{2, 5, 3, 1, 10, 4}
	n := 6
	buildHeap(tree, n)
	for i := 0; i < n; i++ {
		fmt.Println(tree[i])
	}

	// heapsort(&tree, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Println(tree[i])
	// }
}
