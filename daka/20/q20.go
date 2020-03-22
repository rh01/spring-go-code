package main

import "fmt"

func isValid(s string) bool {
	// ret := false
	stack := make([]byte, 0)
	ss := []byte(s)
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, val := range ss {
		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else {
			if len(stack) != 0 && stack[len(stack)-1] == m[val] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
func main() {
	fmt.Println(isValid("(]"))
}
