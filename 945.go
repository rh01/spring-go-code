package main

import (
	"fmt"
	"sort"
)

func minIncrementForUniqueBrute(A []int) int {
	m := make(map[int]int, 0)
	res := 0
	for _, v := range A {
		m[v]++
	}

	for _, v := range A {
		if times, ok := m[v]; ok && times > 1 {
			k := v
			for {
				k++
				res++
				if _, ok := m[k]; !ok {
					m[k]++
					break
				}
			}
			m[v]--
		}
	}
	return res
}

func minIncrementForUniqueOLogN(A []int) int {
	o := 0
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		a, ap := A[i], A[i-1]
		if a <= ap {
			o += ap - a + 1
			A[i] = ap + 1
		}
	}
	return o
}

func minIncrementForUnique(A []int) int {
	pos := make([]int, 80000)
	for i := 0; i < 80000; i++ {
		pos[i] = -1
	}
	move := 0

	var findPos func(int) int
	findPos = func(a int) int {
		b := pos[a]
		// 如果对应的位置为空，直接放入
		if b == -1 {
			pos[a] = a;
			return a
		}

		// 否则向后寻址
		// 由于pos[a] 中标记了上次寻址的空位，因此从pos[a] + 1 开始寻址即可
		b = findPos(b + 1)
		pos[a] = b 
		return b
	}

	for _, a := range A {
		b := findPos(a)
		move += b  - a
	}
	return move
}

func main() {
	ar := []int{3, 2, 1, 2, 1, 7}
	fmt.Println(minIncrementForUnique(ar))
}
