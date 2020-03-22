package main

import "fmt"

func coinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 {
		return -1
	}

	N := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		/// oj := 0
		m := 1<<32 - 1
		for j := 0; j < N; j++ {
			if i >= coins[j] {
				t := dp[i-coins[j]] + 1
				if t < m {
					m = t
				}
			}
			continue
		}
		dp[i] = m
		fmt.Println(dp)
	}
	if dp[amount] == 1<<32-1 {
		return -1
	}
	return dp[amount]
}

func main() {
	// fmt.Println(coinChange([]int{1, 2, 5}, 11))
	// fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{474, 83, 404, 3}, 264))
}
