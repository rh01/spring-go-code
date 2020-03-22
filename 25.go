/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1 
    }
	virtualHead := &ListNode{-1, l1}
	cur := virtualHead
	for {
		if cur.Next.Val > l2.Val {
			tmpNode := l2.Next
			l2.Next = cur.Next
			cur.Next = l2
			l2 = tmpNode	
		}
		cur = cur.Next
        if cur.Next == nil {
            cur.Next = l2
            break
        }
		if l2 == nil {
			break
		}
	}
	return virtualHead.Next
}