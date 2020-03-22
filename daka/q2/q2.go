package main

/**
 *  Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var prev = &ListNode{0, nil}
	walk(prev, head)
	return prev.Next
}

func walk(part1 *ListNode, rem *ListNode) {
	// 这个是边界条件
	if rem == nil {
		return
	}

	rem, rem.Next, part1.Next = rem.Next, part1.Next, rem
	walk(part1, rem)
}

func main() {

}
