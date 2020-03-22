package main

func findSubstring(s string, words []string) []int {
	var res = []int{}
	if len(words) == 0 {
		return res
	}

	// lw 表示word的大小， lwords 表示 words 的長度
	lw, lwords := len(words[0]), len(words)
	// lws 表示words所有word的長度和
	lws := lw * lwords
	// ls 表示s 的長度
	ls := len(s)
	var m = make(map[string]int)
	for i := 0; i < lwords; i++ {
		m[words[i]]++
	}

	for i := 0; i < ls-lws+1; i++ {
		if findIndex(s[i:i+lws], lw, m) {
			res = append(res, i)
		}
	}

	return res
}

func findIndex(s string, lw int, m map[string]int) bool {
	var dict = map[string]int{}
	for i := 0; i < len(s)-lw+1; i++ {
		ss := s[i : i+lw]
		if _, ok := m[ss]; ok {
			dict[ss]++
			i = i + lw - 1
		}
	}

	for k, v := range m {
		if dict[k] != v {
			return false
		}
	}

	return true
}

func main() {

}
