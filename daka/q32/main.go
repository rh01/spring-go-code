package main

func longestValidParentheses(s string) int {
	ret, l, stack := 0, len(s), []int{-1} // 分别表示返回的值，长度和辅助栈存放索引

	for i := 0; i < l; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if 0 == len(stack) {
				stack = append(stack, i)
			} else if t := i - stack[len(stack)-1]; t > ret {
				ret = t
			}
		}
	}
	return ret
}

func main() {

}
