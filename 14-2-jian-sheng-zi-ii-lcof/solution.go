package lcof

func cuttingRopev1(n int) int {
	// 最多可以切n段，由于dp[0]不使用
	dp := make([]int, n+1)
	dp[1] = 1
	res := 0
	for i := 2; i < n+1; i++ {
		j, k := 1, i-1
		for j <= k {
			res = max(res, max(j, dp[j])*max(k, dp[k]))
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func cuttingRope(n int) int {
	// 另一种方法
	maxSum := 1
	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 2
	}

	for n > 4 {
		n = n - 3
		maxSum = maxSum * 3
		if maxSum > 1000000007 {
			maxSum %= 1000000007
		}
	}
	maxSum %= 1000000007
	return maxSum
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
