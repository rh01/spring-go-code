package main

import "fmt"

func main() {
	value := []int{5, 1, 8, 4, 6, 3, 2, 4}

	taskTime := [][2]int{
		[2]int{1, 4},
		[2]int{3, 5},
		[2]int{0, 6},
		[2]int{4, 7},
		[2]int{3, 8},
		[2]int{5, 9},
		[2]int{6, 10},
		[2]int{8, 11},
	}

	prev := computePrevArray(taskTime)
	fmt.Println(prev)

	size := len(value)
	dp := make([]int, size+1)
	dp[0] = 0
	for i := 1; i <= size; i++ {
		dp[i] = max(value[i-1]+dp[prev[i-1]], dp[i-1])
	}
	fmt.Println(dp[size])
}

// 根据taskTime 计算出prev数组
func computePrevArray(taskTime [][2]int) []int {
	size := len(taskTime)
	prev := make([]int, size)

	prev[0] = 0
	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			if taskTime[i][0] >= taskTime[j][1] {
				prev[i] = j + 1
				break
			}
		}
	}
	fmt.Println(prev)
	return prev
}

// max return max value
func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
