package main

func merge(A []int, m int, B []int, n int) {
	var i = m - 1
	var j = n - 1
	var k = m + n - 1

	for i >= 0 && j >= 0 {
		if A[i] >= B[j] {
			A[k] = A[i]
			k--
			i--
		} else {
			A[k] = B[j]
			k--
			j--
		}
	}

	if j >= 0 {
		copy(A[:k+1], B[:j+1])
	}
}

func main() {

}
