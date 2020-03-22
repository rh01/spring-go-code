package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	prefix := prefixTable(str1, str2)
	res := []string{}
	// 计算出公共前缀，然后就可以遍历所有的str1 和 str2
	for i := 0; i < len(prefix); i++ {
		m := len(str1)
		n := len(str2)

		l := len(prefix[i])

		if m%l!= 0 || n%l != 0 {
			continue
		}

		f1, f2 := true, true
		for j := l; j+l <= m; j += l {
			fmt.Println(str1[j : j+l])
			if str1[j:j+l] != prefix[i] {
				f1 = false
			}
		}

		for k := l; k+l <= n; k += l {
			if str2[k:k+l] != prefix[i] {
				f2 = false
			}
		}

		if f1 && f2 {
			res = append(res, prefix[i])
		}
	}

	if len(res) == 0 {
		return ""
	}
	return res[len(res)-1]
}

func prefixTable(str1 string, str2 string) []string {
	res := []string{}
	m, n := len(str1), len(str2)
	i := 0
	for i < min(m, n) {
		if str1[i] == str2[i] {
			res = append(res, str1[:i+1])
		} else {
			break
		}
		i++
	}
	return res
}

func min(i0, i1 int) int {
	if i0 < i1 {
		return i0
	}
	return i1
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
