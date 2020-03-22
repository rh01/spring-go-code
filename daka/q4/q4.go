package main

/**
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。
编写一个方法，将 B 合并入 A 并排序。

初始化 A 和 B 的元素数量分别为 m 和 n。

*/
func merge(A []int, m int, B []int, n int) {
	var i, j, k int
	var tmp = make([]int, m+n)

	for i < m && j < n {
		if A[i] <= B[j] {
			tmp[k] = A[i]
			k++
			i++
		} else {
			tmp[k] = B[j]
			k++
			j++
		}
	}

	if i < m {
		copy(tmp[k:], A[i:m])
	}

	if j < n {
		copy(tmp[k:], B[j:n])
	}

	copy(A, tmp)
}

func main() {

}
