package main

func containsDuplicate(nums []int) bool {
	res := num[0]
	for _, v := range nums[1:] {
		res = res ^ v
		if res == 0 {
			return true
		}
	}

	return false
}

func main() {
	
}