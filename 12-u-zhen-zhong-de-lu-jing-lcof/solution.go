package lcof

// 判断在一个矩阵中是否存在一条包含某字符串所有字符的路径
func exist(board [][]byte, word string) bool {
	// 前提条件
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	// 接下来是用来寻找出口
	rows, cols := len(board), len(board[0])
	flag := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		flag[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if judge(board, i, j, word, 0, flag) {
				return true
			}
		}
	}
	return false
}

func judge(board [][]byte, x, y int, word string, cur int, flag [][]bool) bool {
	// 递归的终止条件
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || flag[x][y] || word[cur] != board[x][y] {
		return false
	}

	if cur == len(word)-1 {
		return true
	}
	flag[x][y] = true

	if judge(board, x+1, y, word, cur+1, flag) || judge(board, x-1, y, word, cur+1, flag) ||
		judge(board, x, y+1, word, cur+1, flag) || judge(board, x, y-1, word, cur+1, flag) {
		return true
	}

	flag[x][y] = false

	return false
}
