package lcof

// CQueue 使用两个栈来实现队列
type CQueue struct {
	stack1 []int // 这个用来模拟队尾
	stack2 []int // 这个用来模拟对头
	// 还可以在这个结构体中添加锁
}

// Constructor 构造函数
func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

// AppendTail 添加到队尾
func (cq *CQueue) AppendTail(value int) {
	cq.stack1 = append(cq.stack1, value)
}

// DeleteHead 对头删除一个元素
func (cq *CQueue) DeleteHead() int {
	if len(cq.stack2) == 0 {
		for len(cq.stack1) != 0 {
			cq.stack2 = append(cq.stack2, cq.stack1[0])
			cq.stack1 = cq.stack1[1:]
		}
	}

	var ret = -1
	if len(cq.stack2) != 0 {
		ret = cq.stack2[0]
		cq.stack2 = cq.stack2[1:]
	}
	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
