func maxProfit(k int, prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	if k >= len(prices) {
		out := 0
		for i := 1; i < len(prices); i++ {
			if prices[i]-prices[i-1] > 0 {
				out += prices[i] - prices[i-1]
			}
		}
		return out
	}
	mr, mc := make([]int, k+1), make([]int, k+1)
	for i := 0; i <= k; i++ {
		mr[i] = -prices[0]
		mc[i] = 0
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			mr[j] = max(mr[j], mc[j-1]-prices[i])
			mc[j] = max(mc[j], mr[j]+prices[i])
		}
	}
	return max(mc[k], 0)
}

func max(i0, i1 int) int {
	if i0 > i1 {
		return i0
	}
	return i1
}
