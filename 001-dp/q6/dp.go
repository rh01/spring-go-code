package main

import "fmt"

func maxProfit(prices []int) int {
	size := len(prices)
	s0 := -prices[0]
	s1 := -1
	for i := 1; i < size; i++ {
		s0 = max(s0, -prices[i])
		s1 = max(s1, s0+prices[i])
	}

	return max(0, s1)
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func main() {
	arr := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(arr))
}
