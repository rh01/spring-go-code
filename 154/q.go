func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		if nums[l] <= nums[r] {
			return nums[l]
		}

		if r == l + 1 {
			return nums[l]
		}

		mid := l + (r -l) >> 1
		if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid
		}
	}
	return -1
}