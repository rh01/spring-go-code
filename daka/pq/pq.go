package main

import "fmt"

type pqElemType = [2]int // 分别是值和索引，索引用来判断是否越界
type pqElemCmpCb = func(pqElemType, pqElemType) bool

// PriorityQueue 优先级队列
type PriorityQueue struct {
	arr    []pqElemType
	lessCb pqElemCmpCb
}

func (pq *PriorityQueue) down(i0, n int) {
	if i0 >= n {
		return
	}

	l, r, m := 2*i0+1, 2*i0+2, i0

	if l < n && pq.lessCb(pq.arr[m], pq.arr[l]) {
		m = l
	}

	if r < n && pq.lessCb(pq.arr[m], pq.arr[r]) {
		m = r
	}

	if m != i0 {
		pq.arr[m], pq.arr[i0] = pq.arr[i0], pq.arr[m]
		pq.down(m, n)
	}
}

func (pq *PriorityQueue) up(j int) {
	if j == 0 {
		return
	}

	p := (j - 1) / 2
	if p != j && pq.lessCb(pq.arr[p], pq.arr[j]) {
		pq.arr[p], pq.arr[j] = pq.arr[j], pq.arr[p]
		pq.up(p)
	}
}

// Init 初始化优先级队列
func (pq *PriorityQueue) Init(arr []pqElemType, lessCb pqElemCmpCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	l := len(pq.arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

// Top 向优先级队列中取最顶上的哪个
func (pq *PriorityQueue) Top() *pqElemType {
	if len(pq.arr) != 0 {
		ret := pq.arr[0]
		return &ret
	}

	return nil
}

// Push 想队列中添加元素
func (pq *PriorityQueue) Push(elem pqElemType) {
	pq.arr = append(pq.arr, elem)
	pq.up(len(pq.arr) - 1)
}

// Pop 向队列中删除一个元素
func (pq *PriorityQueue) Pop() *pqElemType {
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
