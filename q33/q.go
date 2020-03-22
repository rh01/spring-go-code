func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	return search2(nums, 0, len(nums) - 1, target)
}

func search2(nums []int, l, r int, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r - l) >> 1
	if nums[mid] == target {
		return mid
	} 

	if nums[mid] < nums[r] {
		if nums[mid] < target && target <= nums[r] {
			return search2(nums, mid+1, r, target)
		} else {
			return search2(nums, l, mid-1, target)
		}
	} else {
		if nums[l] <= target && target <= nums[mid] {
			return search2(nums, l, mid - 1, target)
		} else {
			return search2(nums, mid + 1, r, target)
		}
	}
}