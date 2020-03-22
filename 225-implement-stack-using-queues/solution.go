package solution

import "container/list"

// MyStack 使用队列实现栈结构
type MyStack struct {
	q1  *list.List // 队列
	q2  *list.List
	top int // top指针
}

// Constructor Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: list.New(),
		q2: list.New(),
	}
}

// Push Push element x onto stack.
func (this *MyStack) Push(x int) {
	this.q1.PushBack(x)
}

// Pop Removes the element on top of the stack and returns that element.
func (this *MyStack) Pop() int {
	for this.q1.Len() > 1 {
		front := this.q1.Front()
		ret := front.Value.(int)
		this.q2.PushBack(ret)
		this.q1.Remove(front)
	}
	poped := this.q1.Remove(this.q1.Front())

	this.q1 = this.q2
	this.q2 = list.New()

	return poped.(int)
}

// Top Get the top element.
func (this *MyStack) Top() int {
	if !this.Empty() {
		return this.q1.Back().Value.(int)
	}
	return -1
}

// Empty Returns whether the stack is empty.
func (this *MyStack) Empty() bool {
	return this.top == 0 || this.q1 == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
