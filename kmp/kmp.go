package main

import "fmt"

// 计算出前缀表
func prefixTable(pattern string, n int) []int {
	var prefix = make([]int, n)
	prefix[0] = 0
	l := 0
	i := 1
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
	i := n - 1
	for ; i > 0; i-- {
		prefix[i] = prefix[i-1]
	}
	prefix[0] = -1
}

func kmpSearch(pattern string, t string) {
	m := len(t)
	n := len(pattern)
	prefix := prefixTable(pattern, n)
	movePrefixTable(prefix, n)

	// text[i]   	len(text)    = m
	// pattern[j]	len(pattern) = n
	i := 0
	j := 0
	for i < m {
		if j == n-1 && t[i] == pattern[j] {
			fmt.Printf("Found Pattern  at %d\n", i-j)
			j = prefix[j]
		}
		if t[i] == pattern[j] {
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
	pattern := "ABABCABAA"
	prefix := prefixTable(pattern, len(pattern))
	movePrefixTable(prefix, len(prefix))
	fmt.Println(prefix)

	kmpSearch(pattern, "ABABABCABAA")
}
