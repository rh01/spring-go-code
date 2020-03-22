package main

func checkInclusion(s1 string, s2 string) bool {
	// 这里根据提示选择切片来记录s1中每个字符的出现的频率
	var arr = make([]int, 128)
	for _, s := range s1 {
		arr[s]++
	}

	// 接下来记录最左边和最右边的index，用来判断长度是否等于s1的长度
	l, r, cnt := 0, 0, len(s1)
	for r = 0; r < len(s2); r++ {
		if arr[s2[r]] > 0 {
			cnt--
		}
		arr[s2[r]]--

		for cnt == 0 {
			if r-l+1 == len(s1) {
				return true
			}

			arr[s2[r]]++
			if arr[s2[r]] > 0 {
				cnt++
			}
			l++
		}
	}
	return false
}

func main() {

}
