package lcof

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)

	// 初始化
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 特判 空正则
			if j == 0 {
				dp[i][j] = i == 0
			} else {
				// 非空正则,分两种情况，* 和非*
				if p[j-1] != '*' {
					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
						dp[i][j] = dp[i-1][j-1]
					}
				} else {
					// 遇到 * 了，分为看和不看两种情况
					// 不看
					if j >= 2 {
						dp[i][j] = dp[i][j] || dp[i][j-2]
					}
					// 看
					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						dp[i][j] = dp[i][j] || dp[i-1][j]
					}
				}
			}
		}
	}
	return dp[n][m]

}

/**
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符
，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。

例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

*/
