package main

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	dp := make([]bool, size+1)

	for i := 1; i <= size; i++ {
		for j := i - 1; j >= 0; j-- {
			if find(s[i:j], wordDict) && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[size]
}

func find(s string, wordDict []string) bool {
	for _, v := range wordDict {
		if v == s {
			return true
		}
	}
	return false
}

func main() {

}
