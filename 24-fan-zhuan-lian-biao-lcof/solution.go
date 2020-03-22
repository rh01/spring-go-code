package lcof

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 引入虚头结点
	vhead := &ListNode{Next: nil}
	cur := head
	for cur != nil {
		// temp := cur.Next
		// cur.Next = vhead.Next
		// vhead.Next = cur
		// cur = temp
		// 转换一下
		cur.Next, cur, vhead.Next = vhead.Next, cur.Next, cur
	}
	return vhead.Next
}
