package main

import "fmt"

// 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

// 示例 1:

// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// 示例 2:

// 输入: [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

func maxProduct(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	// dp数组的第一个元素是自身
	// dp[0] = nums[0]
	var minV int
	for i := 0; i < size; i++ {
		dp[i] = nums[i]
		minV = nums[0]
	}

	for i := 1; i < size; i++ {
		if nums[i] >= 0 {
			dp[i] = max(nums[i], nums[i]*dp[i-1])
			minV *= nums[i]
		}

		fmt.Println(minV)

		if nums[i] < 0 {
			dp[i] = max(dp[i], minV*nums[i])
			minV = min(nums[i], nums[i]*dp[i-1])
		}
	}

	

	var maxV = dp[0]
	for i := 1; i < size; i++ {
		if dp[i] > maxV {
			maxV = dp[i]
		}
	}
	fmt.Println(dp)
	return maxV
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func main() {
	arr := []int{2, -5, -2, -4, 3}
	fmt.Println(maxProduct(arr))
}
