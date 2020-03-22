package main

import "fmt"

func PathNum(arr [][]byte) int {
	if arr == nil || len(arr) == 0 || len(arr[0]) == 0 {
		return -1
	}

	m, n := len(arr), len(arr[0])
	fmt.Println(m, n)
	nums := make([][]int, m)

	for i := 0; i < m; i++ {
		nums[i] = make([]int, n)
	}

	nums[0][0] = 1

	for i := 1; i < n; i++ {
		if arr[0][i] == '.' {
			nums[0][i] = nums[0][i-1]
		}
	}

	if nums[0][1] == 1 {
		nums[1][0] = 1
	}

	for i := 1; i < n; i++ {
		if arr[1][i] == '.' {
			nums[1][i] = nums[1][i-1]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums[i][j] != '.' {
				nums[i][j] = nums[i][j] + nums[i-1][j] + nums[i][j-1] + nums[i-1][j] 
			}
		}
	}

	return nums[m-1][n-1]
}

func main() {
	arr := [][]byte{
		[]byte{'.', '.', 'X', '.', 'X'},
		[]byte{'X', 'X', '.', '.', '.'},
	}

	fmt.Println(PathNum(arr))
}
