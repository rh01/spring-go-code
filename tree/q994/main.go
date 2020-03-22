package main

var dirs = [4][2]int8{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func orangesRotting(grid [][]int) int {

	// o 表示xxx, m 代表行数， n 代表列数
	o, m, n := 0, uint8(len(grid)), uint8(len(grid[0]))
	// q表示队列，记录已经坏的橘子，mk表示字典，记录是否访问
	q, mk := []uint8{}, map[uint8]int{}

	for i := uint8(0); i < m; i++ {
		for j := uint8(0); j < n; j++ {
			if grid[i][j] == 2 {
				k := i<<4 | j
				q, mk[k] = append(q, k), 0
			}
		}
	}

	for len(q) != 0 {
		k := q[0] // 出队
		q = q[1:]

		// 然后通过二进制再将它们计算出来
		var y, x uint8 = k >> 4, k & 0x0f
		for _, d := range dirs {
			i, j := int8(y)+d[0], int8(x)+d[1]
			//  出界了
			if i < 0 || i >= int8(m) || j < 0 || j >= int8(n) || grid[i][j] != 1 {
				continue
			}

			// 如果非空，并且也不是已经坏的橘子，则将其
			grid[i][j] = 2

			// 然后再将坏的橘子，重新入队列
			k1 := uint8(i<<4 | j)
			q, mk[k1] = append(q, k1), mk[k]+1
			o = mk[k1]
		}
	}

	for i := uint8(0); i < m; i++ {
		for j := uint8(0); i < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return o
}

func main() {

}
