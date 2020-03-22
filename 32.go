package main

import "fmt"

func prefixTable(pattern string, n int) []int {
	prefix := make([]int, n)
	prefix[0] = 0
	i := 1
	l := 0
	for i < n {
		if pattern[l] == pattern[i] {
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
	var res = make([]int, n)
	res[0] = -1
	copy(res[1:], prefix[:n-1])
	return res
}

func kmpSearch(text string, pattern string) {
	m, n := len(text), len(pattern)
	prefix := prefixTable(pattern, n)

	i, j := 0, 0
	for i < m {
		if j == n-1 && pattern[j] == text[i] {
			fmt.Println("found at", i-n+1)
		}
		if text[i] == pattern[j] {
			i++
			j++
		} else {
			j = prefix[j]
			if j == -1 {
				j++
				i++
			}
		}
	}
}

func main() {
	text := "ABABABC"
	pattern := "ABC"

	kmpSearch(text, pattern)
}
