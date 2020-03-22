package lcof

func hammingWeight(n int) int {
	var res int
	for n != 0 {
		n = n & (n - 1)
		res++
	}
	return res
}
