package main

func findDuplicate(nums []int) int {
	if 0 == len(nums) {
		return -1
	}

	slow, fast := 0, 0
	for {
		// 这个可以看做是快慢指针
		slow, fast = nums[slow], nums[nums[fast]]
		// 找到环
		if slow == fast {
			break
		}
	}

	// 接下来开始一步一步走
	fast = nums[0]
	for fast != slow {
		fast  = nums[fast]
		slow = nums[slow]
	}
	return nums[slow]
}

func main() {

}
