package main

func heapify(arr []int, i int, n int) {
	if i >= n {
		return
	}

	l, r, m := 2*i+1, 2*i+2, i
	if l < n && arr[l] < arr[m] {
		m = l
	}

	if r < n && arr[l] < arr[m] {
		m = r
	}

	if m != i {
		arr[m], arr[i] = arr[i], arr[m]
		heapify(arr, m, n)
	}
}

func buildHeap(arr []int, n int) {
	l := n>>1 - 1
	for ; l >= 0; l-- {
		heapify(arr, l, n)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	buildHeap(arr, n)
	res := []int{}
	for i := 1; i <= k; i++ {
		arr[0], arr[n-i] = arr[n-i], arr[0]
		res = append(res, arr[n-i])
		heapify(arr, 0, n-i)
	}
	return res
}

func main() {

}
