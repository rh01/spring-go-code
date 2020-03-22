package substring

func longestPalindrome(s string) string {

	maxLength := 0 // 记录向两边扩展的长度
	//idx := -1      // 记录中心的索引
	var start, end int
	for i := 0; i < len(s); i++ {
		l1 := extendFromCenter(s, i, i)
		l2 := extendFromCenter(s, i, i+1)
		if max(l1, l2) > maxLength {
			maxLength = max(l1, l2)
			//idx = i
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}
	}
	// 确定开始索引和结束索引

	return s[start : end+1]
}

// 这个是有start-end向两边开始扩展
func extendFromCenter(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return end - start - 1
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
