package lcof

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// 算法的逻辑从左下角开始寻找或者右上角开始找，即可
	rows, cols := len(matrix)-1, len(matrix[0])-1
	// 从左下角开始的逻辑
	// 实践证明从左下角开始的逻辑是错的
	r := 0
	c := cols
	for r <= rows && c >= 0 {
		if matrix[r][c] > target {
			c--
		} else if matrix[r][c] < target {
			r++
		} else {
			return true
		}
	}
	return false
}
