package matching

func isMatch(s string, p string) bool {
	l1 := len(s)
	l2 := len(p)

	dp := make([][]bool, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]bool, l2)
	}

	dp[0][0] = true
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			// 分两种情况
			if p[j-1] != '*' {
				if s[i-1] == p[j-1] || p[j-1] == '.' {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if j >= 2 {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}

				if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	return dp[l1][l2]
}
