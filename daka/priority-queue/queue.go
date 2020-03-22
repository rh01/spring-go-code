package main

import "fmt"

type pqElemeType = [2]int
type pqElemCmpCb = func(pqElemeType, pqElemeType) bool

type PriorityQueue struct {
	arr    []pqElemeType
	lessCb pqElemCmpCb
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2

		if i == j || !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}

		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		j = i
	}
}

func (pq *PriorityQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := i<<1 + 1
		if j1 >= n || j1 < 0 {
			break
		}

		j := j1
		if j2 := j1 + 1; j2 < n && pq.lessCb(pq.arr[j2], pq.arr[j1]) {
			j = j2
		}

		if !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr []pqElemeType, lessCb pqElemCmpCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	l := len(pq.arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueue) Top() *pqElemeType {
	if len(pq.arr) != 0 {
		e := pq.arr[0]
		return &e
	}
	return nil
}

func (pq *PriorityQueue) Push(item pqElemeType) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() *pqElemeType {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		e := pq.arr[i]
		pq.arr = pq.arr[:i]
		return &e
	}

	return nil
}

func maxSlidingWindow(nums []int, k int) []int {
	out := []int{}
	//  定义一个优先级队列
	q := (&PriorityQueue{}).Init(nil, func(a, b [2]int) bool {
		return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
	})

	for i := range nums {
		for {
			//  0 - k 并且还要判断范围 直接break，
			if p := q.Top(); p == nil || (*p)[1] > i-k {
				break
			}
			q.Pop()
		}

		q.Push([2]int{nums[i], i})
		// 0-k下面不执行，从现在开始可以
		if i >= k-1 {
			p := q.Top()
			out = append(out, (*p)[0])
		}
	}
	return out
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
