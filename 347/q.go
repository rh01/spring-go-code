package main

type elemType = [2]int
type elemCmp = func(elemType, elemType) bool

// PriorityQueue 优先级队列
type PriorityQueue struct {
	arr     []elemType
	lessCmp elemCmp
}

func (pq *PriorityQueue) up(j int) {
	if j == 0 {
		return
	}

	p := (j - 1) / 2
	if j != p && pq.lessCmp(pq.arr[p], pq.arr[j]) {
		pq.arr[p], pq.arr[j] = pq.arr[j], pq.arr[p]
		pq.up(p)
	}
}

func (pq *PriorityQueue) down(i0 int, n int) {
	if i0 >= n {
		return
	}

	l, r, m := 2*i0+1, 2*i0+2, i0

	if l < n && pq.lessCmp(pq.arr[m], pq.arr[l]) {
		m = l
	}

	if r < n && pq.lessCmp(pq.arr[m], pq.arr[r]) {
		m = r
	}

	if m != i0 {
		pq.arr[i0], pq.arr[m] = pq.arr[m], pq.arr[i0]
		pq.down(m, n)
	}
}

// Init 初始化New, Push, Pop
func (pq *PriorityQueue) Init(arr []elemType, lessCmp elemCmp) *PriorityQueue {
	pq.arr = arr
	pq.lessCmp = lessCmp
	n := len(arr)
	for i := n>>1 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
	return pq
}

// Push 添加元素
func (pq *PriorityQueue) Push(item elemType) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

// Pop 删除一个元素
func (pq *PriorityQueue) Pop() *elemType {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		item := pq.arr[i]
		pq.arr = pq.arr[:i]
		pq.down(0, i)
		return &item
	}
	return nil
}

func topKFrequent(nums []int, k int) []int {
	var m = map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	q := (&PriorityQueue{}).Init(nil, func(a, b [2]int) bool {
		return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
	})

	o := []int{}
	for n, c := range m {
		q.Push([2]int{c, n})
	}

	for i := 0; i < k; i++ {
		o = append(o, (*q).Pop()[1])
	}

	return o
}
