package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	r := target / 2
	if target%2 == 0 {
		r = target/2 + 1
	}
	res := make([][]int, 0)
	for i := 1; i <= r; i++ {
		for j := i + 1; j <= r+1; j++ {
			if (j-i+1)*(i+j)/2 == target {
				res = append(res, []int{i, j})
				break
			}
		}
	}

	return res
}

func main() {
	target := 15
	fmt.Println(findContinuousSequence(target))
}
