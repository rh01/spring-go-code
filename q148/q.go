/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, s, f := head, head, head
	for nil != f && nil != f.Next {
		p, s, f = s, s.Next, f.Next.Next
	}
	p.Next = nil

	return merge(sortList(head), sortList(s))
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l2.Next, l1)
		return l2
	}
}
