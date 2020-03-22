package main

func buildHeap(nums []int, n int) {
	// buildHeap的過程，從最後一個元素開始
	l, p := n-1, (l-1)/2
	for i := p; i >= 0; i-- {
		heapify(nums, n, i)
	}

}

func heapify(nums []int, n int, i int) {
	// 边界条件
	if i >= n {
		return
	}

	l, r := 2*i+1, 2*i+2

	m := i

	if l < n && nums[m] < nums[l] {
		m = l
	}

	if r < n && nums[m] < nums[r] {
		m = r
	}

	if m != n {
		nums[i], nums[m] = nums[m], nums[i]
		heapify(nums, n, m)
	}
}

// func heapSort(nums []int, n int) {

// }

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n < k {
		return -1
	}

	var tmp = make([]int, n)
	copy(tmp, nums)
	buildHeap(tmp, n-k)
	for i := n - k; i < n; i++ {

	}

}

func main() {

}
