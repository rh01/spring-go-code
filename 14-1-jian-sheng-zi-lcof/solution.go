package lcof

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	// 首项
	dp[1] = 1
	// 然后考虑如何切
	for i := 2; i < n+1; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max(res, max(j, dp[j])*max(k, dp[k])) // 递推公式
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max(i1, i2 int) int {
	if i1 >= i2 {
		return i1
	}
	return i2
}
