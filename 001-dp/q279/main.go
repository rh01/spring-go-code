package main

func numSquares(n int) int {
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 0
		dp[0][i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j*j <= i {
				dp[i][j] = dp[i-j*j][j] + 1
			}
		}
	}
}

func main() {

}
