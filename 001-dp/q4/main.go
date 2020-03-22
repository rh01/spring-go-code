package main

import "fmt"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp := make([]int, size)

	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < size; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(arr))
}
