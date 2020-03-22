package main

import "fmt"

func dp(arr []int) int {
	size := len(arr)
	dp := make([]int, size)
	dp[0] = arr[0]
	// dp[1] = arr[0]
	for i := 1; i < size; i++ {
		if i == 1 {
			dp[i] = max(arr[i], dp[i-1])
		} else {
			dp[i] = max(arr[i]+dp[i-2], dp[i-1])
		}
	}
	return dp[size-1]
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func main() {

	arr := []int{1, 2, 4, 1, 7, 8, 3}
	fmt.Println(dp(arr))
}
