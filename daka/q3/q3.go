package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var vHead = &ListNode{0, head}
	var pre, next *ListNode
	cur := vHead
	count := 1

	for ; cur != nil && count != m; count++ {
		cur = cur.Next
	}
	pre = cur

	for ; cur != nil && count != n+1; count++ {
		cur = cur.Next
	}
	next = cur
	after := next.Next
	next.Next = nil
	mid := pre.Next
	pre.Next = nil
	pre.Next = after

	cur = mid
	for cur != nil {
		cur, cur.Next, pre.Next = cur.Next, pre.Next, cur
	}
	return vHead.Next
}

func main() {
}
