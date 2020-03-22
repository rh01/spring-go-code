func maxProfit(prices []int) int {
	i0, i1 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		i0 = max(i0, -prices[i])
		i1 = max(i1, i0+prices[i])
	}
	return max(i1, 0)
}

func max(i0, i1 int) int {
	if i0 > i1 {
		return i0
	}
	return i1
}
