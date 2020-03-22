package main

import "fmt"

func qsort(x []int, l, u int) {
	if l >= u {
		return
	}

	i, j := l+1, u
	t := x[l]
	for {

		for i <= u && x[i] < t {
			i++
		}

		for x[j] > t {
			j--
		}

		if i > j {
			break
		}

		x[i], x[j] = x[j], x[i]
	}

	x[l], x[j] = x[j], x[l]

	qsort(x, l, j-1)
	qsort(x, j+1, u)

}

func main() {
	x := []int{2, 3, 1, 4, 5}
	qsort(x, 0, len(x)-1)
	fmt.Println(x)
}
