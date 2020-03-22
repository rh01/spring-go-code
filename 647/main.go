package main

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += count(s, i, i)
		res += count(s, i, i + 1)
	}
	return res
}

func count(s string, l, r int) int {
	cnt := 0
	for l >=0 && l < len(s) && s[l] == s[r] {
		cnt++
		l--
		r++
	}
	return cnt
}



func main() {
	
}