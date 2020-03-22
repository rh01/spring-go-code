package main

func main() {

}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := &ListNode{0, nil}
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = tmp
	}
	
	return prev.Next
}
