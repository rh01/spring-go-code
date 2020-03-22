func validateStackSequences(pushed []int, popped []int) bool {
	i, j := 0, 0
	st := []int{}
	for {
		if len(st) != 0 && st[len(st) - 1] == popped[j] {
			st, j = st[:len(st) - 1], j + 1
		} else {
			if i >= len(popped) {
				break
			}
			st, i = append(st, pushed[i]), i + 1
		}
	}
	return 0 == len(st)
}