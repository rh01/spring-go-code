package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	size := len(s)
	if size == 0 || size == 1 {
		return size
	}

	var prev = make([]int, size)
	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				prev[i] = j + 1
				break
			}
		}
		if (prev[i-1] != 0 && prev[i] == 0) || (prev[i-1] > prev[i]) {
			prev[i] = prev[i-1]
		}
	}
	fmt.Println(prev)

	var dp = make([]int, size)
	dp[0] = 1
	for i := 1; i < size; i++ {
		dp[i] = max(dp[i-1], i+1-prev[i])
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
	var s string = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
