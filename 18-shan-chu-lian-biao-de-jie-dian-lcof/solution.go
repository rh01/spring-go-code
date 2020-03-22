package lcof

/**
 * Definition for singly-linked list.
 *  */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	// 如果当前头结点为空或者头结点的next节点为空，那么直接返回nil
	if head == nil || (head.Next == nil && head.Val == val) {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	// 下面寻找与给定val相等的节点，并将其删除掉即可
	// 这里使用两个指针操作
	cur := head
	pre, next := cur, cur.Next
	for next != nil {
		if next.Val == val {
			// tmp := next.Next
			pre.Next = next.Next
			break
		}
		pre = pre.Next
		next = next.Next
	}
	return cur
}

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// 返回删除后的链表的头节点。
// 注意：此题对比原题有改动
/**
func deleteNode(head *ListNode, val int) *ListNode {
	// 如果当前头结点为空或者头结点的next节点为空，那么直接返回nil
	if head == nil  {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	// 下面寻找与给定val相等的节点，并将其删除掉即可
	// 这里使用两个指针操作
	cur := head
	pre, next := cur, cur.Next
	for next != nil && next.Val != val {
        pre = pre.Next
        next = next.Next
	}
    pre.Next = next.Next
	return cur
}*/