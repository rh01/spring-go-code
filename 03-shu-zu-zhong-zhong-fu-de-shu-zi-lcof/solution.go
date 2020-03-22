package lcof

func findRepeatNumberV1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	sz := len(nums)
	hmap := make(map[int]bool, sz)

	for _, num := range nums {
		if _, exist := hmap[num]; !exist {
			hmap[num] = true
		} else {
			return num
		}
	}
	return -1
}

func findRepeatNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	sz := len(nums)
	harray := make([]int, sz)
	for i := 0; i < sz; i++ {
		harray[nums[i]]++
		if harray[nums[i]] > 1 {
			return nums[i]
		}
	}

	return -1
}
