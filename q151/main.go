package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	n := len(arr)

	words := make([]string, 0, n)
	for i := n - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			words = append(words, arr[i])
		}
	}


	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
