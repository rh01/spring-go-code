package lcof

// 双指针
func exchange(nums []int) []int {
	// // // var res = make([]int, 0)
	sz := len(nums)
	lo, hi := 0, sz-1
	for lo < hi {
		for lo < hi && nums[lo]%2 != 0 {
			lo++
		}

		for lo < hi && nums[hi]%2 == 0 {
			hi--
		}

		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
	return nums
}
