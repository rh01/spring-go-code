package main

import "fmt"

func computeDP(arr []int, targetS int) bool {
	size := len(arr)
	// 这里的多分配一个targetS 0
	// const targetS = S
	var dp = make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, targetS+1)
	}

	for i := 0; i < size; i++ {
		dp[i][0] = true // 因为这里大家都可以不选构成0
	}

	for i := 1; i <= targetS; i++ {
		dp[0][i] = arr[0] == i
	}

	for i := 1; i < size; i++ {
		for j := 1; j <= targetS; j++ {
			// 选或不选
			if arr[i] > j {
				dp[i][j] = dp[i-1][j]
			} else if arr[i] == j {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i-1][j-arr[i]] || dp[i-1][j]
			}
		}
	}
	return dp[size-1][targetS]
}

func main() {
	arr := []int{3, 34, 4, 12, 5, 2}
	fmt.Println(computeDP(arr, 33))
}
