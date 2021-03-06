package lcof

func numWays(n int) int {
	f0, f1 := 0, 1

	for n != -1 {
		f0, f1 = f1, f0+f1
		n--
		f0 = f0 % (1e9 + 7)
	}

	return f0
}
