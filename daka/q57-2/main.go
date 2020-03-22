package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	arr := make([]int, (target+1)/2+1)
	for i := 1; i <= (target+1)/2; i++ {
		arr[i] = (i + 1) * i / 2
	}

	m := make(map[int]int, (target+1)/2+1)
	m[0] = 0
	for i := 1; i <= (target+1)/2; i++ {
		m[i*(i+1)/2] = i
	}

	res := make([][]int, 0)
	for i := 1; i <= (target+1)/2; i++ {
		s := i * (i + 1) / 2
		if s >= target {
			if v, exist := m[s-target]; exist {
				k := v + 1
				var tmp = make([]int, 0)
				for k <= i {
					tmp = append(tmp, k)
					k++
				}
				res = append(res, tmp)
			}
		}
	}

	return res
}

func main() {
	target := 15
	fmt.Println(findContinuousSequence(target))
}
