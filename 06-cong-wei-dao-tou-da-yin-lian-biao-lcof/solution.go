package lcof

/*
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 栈实现
func reversePrintV1(head *ListNode) []int {
	var ret = make([]int, 0)
	if head == nil {
		return ret
	}
	cur := head
	for cur != nil {
		ret = append([]int{cur.Val}, ret...)
		cur = cur.Next
	}
	return ret
}

// 递归实现
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint(head.Next), head.Val)
}
