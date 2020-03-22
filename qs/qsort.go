package main

import "fmt"

func qsort(x []int, l, r int) {

	if l >= r {
		return
	}

	m := l
	// val := arr[l]
	for i := l + 1; i < r; i++ {
		if x[i] < x[l] {
			m++
			x[m], x[i] = x[i], x[m]
		}

	}
	x[l], x[m] = x[m], x[l]
	qsort(x, l, m-1)
	qsort(x, m+1, r)
}

func main() {
	x := []int{2, 3, 1, 4, 5}
	qsort(x, 0, len(x)-1)
	fmt.Println(x)
}