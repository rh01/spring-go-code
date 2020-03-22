package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1, l2 := len(nums1), len(nums2)
	nums := make([]int, l1+l2)
	copy(nums[:l1], nums1)
	copy(nums[l1:], nums2)

	// fmt.Println(nums)

	buildHead(nums, l1+l2)
	// fmt.Println(nums)

	// 分两种情况吧
	n := l1 + l2
	var i int
	for i = n - 1; i >= n/2+1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	t1 := nums[0]
	fmt.Println(t1)
	if n%2 != 0 {
		return float64(t1)
	}

	nums[i], nums[0] = nums[0], nums[i]
	heapify(nums, i, 0)
	t2 := nums[0]
	fmt.Println(t2)

	return float64(t1+t2) / 2.0
}

func buildHead(nums []int, n int) {
	l := n - 1 // last node
	p := (l - 1) / 2
	for i := p; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

func heapify(nums []int, n int, i int) {
	if i >= n {
		return
	}

	m, l, r := i, 2*i+1, 2*i+2

	if l < n && nums[l] > nums[m] {
		m = l
	}

	if r < n && nums[r] > nums[m] {
		m = r
	}

	if m != i {
		nums[m], nums[i] = nums[i], nums[m]
		heapify(nums, n, m)
	}

}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
