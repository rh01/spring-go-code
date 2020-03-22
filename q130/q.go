package main

func solve(board [][]byte) {
	if board == nil || len(board) == 0 || board[0] == nil || len(board) == 0 {
		return
	}

	rs, ls := len(board), len(board[0])

	for i := 0; i < rs; i++ {
		dfs(board, rs, ls, i, 0)
		dfs(board, rs, ls, i, ls-1)
	}

	for i := 0; i < ls; i++ {
		dfs(board, rs, ls, 0, i)
		dfs(board, rs, ls, rs-1, i)
	}

	for i := 0; i < rs; i++ {
		for j := 0; j < ls; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'o' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfs(board [][]byte, rs, ls int, x, y int) {
	// stop condition
	if x < 0 || x >= rs || y < 0 || y >= ls || board == 'o' {
		return
	}

	if board[x][y] == 'O' {
		board[x][y] = 'o'
		dfs(board, rs, ls, x+1, y)
		dfs(board, rs, ls, x-1, y)
		dfs(board, rs, ls, x, y+1)
		dfs(board, rs, ls, x, y-1)
	}
}
