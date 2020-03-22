/**
 * Definition for singly-linked list.
 */ 
type ListNode struct {
      Val int
      Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertionSortList(head *ListNode) *ListNode {
	// 当头结点或者下一个节点为空，则返回头结点
	if head == nil || head.Next == nil {
		return head
	}

    cur2 := head.Next
    head.Next = nil

	vHead := &ListNode{-1, head}
    cur := vHead	

	for cur2 != nil {
		tmp :=cur2.Next
        var m = cur
		for ; m.Next != nil && m.Next.Val < cur2.Val; m = m.Next {
		} 
		cur2.Next = m.Next
		m.Next = cur2
		cur2 = tmp
	}
	return vHead.Next
}