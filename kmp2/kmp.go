package main

import "fmt"

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
	l := n - 1
	for ; l > 0; l-- {
		prefix[l] = prefix[l-1]
	}
	prefix[0] = -1
}

func kmpSearch(text string, pattern string) {

	// text    length = n, i
	// pattern length = m, j
	n := len(text)
	m := len(pattern)
	i := 0
	j := 0

	prefix := prefixTable(pattern, m)
	movePrefixTable(prefix, m)

	for i < n {
		if j == m-1 && text[i] == pattern[j] {
			fmt.Printf("Found pattern at %d\n", i-j)
			j = prefix[j]
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
	pattern := "ABABAAB"
	//movePrefixTable(prefix, 7)
	//fmt.Println(prefix)
	kmpSearch("ABABABAABBAA", pattern)
}
