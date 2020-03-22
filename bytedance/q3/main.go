package main

import "fmt"

func getMinLengthString(strs []string) (string, int) {

	size := len(strs)
	if size <= 1 {
		return strs[0], len(strs[0])
	}

	minStr, minLen := strs[0], len(strs[0])
	for i := 1; i < size; i++ {
		if len(strs[i]) < minLen {
			// fmt.Println("jinlai")
			minStr = strs[i]
			minLen = len(strs[i])
		}
	}
	return minStr, minLen
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	minStr, minLen := getMinLengthString(strs)
	fmt.Println(minStr, minLen)
	i := 0
	for i < minLen {
		for _, s := range strs {
			if s[i] != minStr[i] {
				return minStr[:i]
			}
		}
		// ret++
		i++
	}
	return minStr
}

func main() {
	a := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(a))
}
