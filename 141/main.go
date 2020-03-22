package main

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow, fast := head.Next, head.Next.Next
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		if fast == slow {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
}

func main() {

}
