package main

import "fmt"

func prefixTable(pattern string, n int) []int {
	prefix := make([]int, n)
	prefix[0] = 0
	l, i := 0, 1
	for i < n {
		if pattern[i] == pattern[l] {
			l++
			prefix[i] = l
			i++
		} else {
			if l > 0 {
				l = prefix[l-1]
			} else {
				prefix[i] = l
				i++
			}
		}
	}
	return prefix
}

func movePrefixTable(prefix []int, n int) {
	for last := n - 1; last > 0; last-- {
		prefix[last] = prefix[last-1]
	}
	prefix[0] = -1
}

func kmpSearch(text string, pattern string) {
	i, j := 0, 0
	m, n := len(text), len(pattern)

	// 假定m大于n
	prefix := prefixTable(pattern, n)
	movePrefixTable(prefix, n)
	for i < m {
		// 前面判断xxxx
		if j == n-1 && pattern[j] == text[i] {
			fmt.Println("found")
		}
		if text[i] == pattern[j] {
			i++
			j++
		} else {
			j = prefix[j]
			if j == -1 {
				i++
				j++
			}
		}
	}
}

func main() {

}
