func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	i0, i1 := -prices[0], 0
	i2, i3 := 0, 0

	for i := 1; i < len(prices); i++ {
		i0 = max(i0, -prices[i])
		i1 = max(i1, i0+prices[i])

		i2 = max(i2, i1-prices[i])
		i3 = max(i3, i2+prices[i])
	}

	return i3
}

func max(i0, i1 int) int {
	if i0 > i1 {
		return i0
	}
	return i1
}
