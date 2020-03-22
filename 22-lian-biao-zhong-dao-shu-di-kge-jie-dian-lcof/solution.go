package lcof

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}

	var count int
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}

	if k > count {
		return nil
	}

	cur = head
	for count-k != 0 {
		cur = cur.Next
		k++
	}
	return cur
}
