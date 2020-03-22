package main

import "fmt"

func countCharacters(words []string, chars string) int {
	if words == nil || len(words) == 0 {
		return 0
	}

	res := 0
	patternM := make(map[byte]int)

	for _, c := range chars {
		patternM[byte(c)]++
	}

	for i := 0; i < len(words); i++ {
		var m = make(map[byte]int)
		for k, v := range patternM {
			m[k] = v
		}
		found := true
		for _, cc := range words[i] {
			n, ok := m[byte(cc)]
			if ok && n > 0 {
				m[byte(cc)]--
			} else {
				found = false
				break
			}
		}
		if found {
			res += len(words[i])
		}
	}

	return res
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))

}
