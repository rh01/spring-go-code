package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	if tokens == nil || len(tokens) == 0 {
		return -1
	}

	var opMap = map[string]func(int, int) int{
		"+": func(i0, i1 int) int { return i0 + i1 },
		"-": func(i0, i1 int) int { return i0 - i1 },
		"*": func(i0, i1 int) int { return i0 * i1 },
		"/": func(i0, i1 int) int { return i0 / i1 },
	}

	var stack = []int{}
	var ok bool
	for _, v := range tokens {

		if _, ok = opMap[v]; ok {
			// fmt.Println("in")
			i1, i0 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			e := opMap[v](i0, i1)
			stack = append(stack, e)
		} else {
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		}
	}
	return stack[0]
}

func main() {
	ops := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(ops))
}
