package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		if fast == slow {
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {

}


