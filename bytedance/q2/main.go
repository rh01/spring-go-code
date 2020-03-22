package main

// 下面实现一种滑动窗口的做法
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size <= 1 {
		return size
	}

	l, r, i, ret := 0, 0, 0, 0 // 分别记录着最左\最右\索引和结果

	for ; r < size; r++ {
		for i = l; i < r; i++ {
			if s[i] == s[r] {
				l = i + 1 // 此时更新最左边的值
				break
			}
		}
		// 这里计算最新的长度
		if r - l + 1 > ret {
			ret = r - l + 1
		}
	}
	return ret
}



func main() {

}
