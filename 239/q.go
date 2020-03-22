package main

import "fmt"

func heapify(nums []int, n int, i int) {
	if i >= n {
		return
	}

	l, r, m := 2*i+1, 2*i+2, i

	if l < n && nums[l] > nums[m] {
		m = l
	}

	if r < n && nums[r] > nums[m] {
		m = r
	}

	if m != i {
		nums[i], nums[m] = nums[m], nums[i]
		heapify(nums, n, m)
	}
}

func buildHeap(nums []int, n int) {
	last := n - 1
	parent := (last - 1) / 2

	for i := parent; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	// fmt.Println(tmp)

	var res []int
	for i := 0; i <= n-k; i++ {
		tmp := make([]int, k)
		copy(tmp, nums[i:i+k])
		fmt.Println(tmp)
		buildHeap(tmp, k)
		res = append(res, tmp[0])
		// tmp[0] = nums[i]
		// heapify(tmp, k, 0)
		// tmp = []int{}
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
