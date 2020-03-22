package lcof

func movingCount(m int, n int, k int) int {
	// 初始化方格系统
	var grid = make([][]bool, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}

	grid[0][0] = false
	return walk(grid, 0, 0, m, n, k)
}

func walk(grid [][]bool, x, y, m, n, k int) int {
	// 停止条件
	if x >= m || x < 0 || y >= n || y < 0 || grid[x][y] || x%10+x/10+y%10+y/10 > k {
		return 0
	}

	grid[x][y] = true
	return walk(grid, x+1, y, m, n, k) + walk(grid, x-1, y, m, n, k) +
		walk(grid, x, y+1, m, n, k) + walk(grid, x, y-1, m, n, k) + 1

}

// 工具函数，用来判断是否到达了行坐标与横坐标的数位之和
// 等于给定的值k，如果到达则返回true否则返回false
func isReachK(x, y, k int) bool {
	if getSum(x)+getSum(y) > k {
		return true
	}
	return false
}

// 该方法用来获取数位之和
func getSum(x int) int {
	var res int
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}
