package lcof

import "testing"

func TestReversePrint(t *testing.T) {
	// 输入：head = [1,3,2]
	// 输出：[2,3,1]
	head := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				2,
				nil,
			},
		},
	}
	expect := []int{2, 3, 1}
	if res := reversePrint(head); !compareSlice(expect, res) {
		t.Errorf("expect %v, but got %v\n", expect, res)
	}
}

func compareSlice(slice1, slice2 []int) bool {
	if (slice1 == nil && slice2 != nil) || (slice1 != nil && slice2 == nil) ||
		len(slice1) != len(slice2) {
		return false
	}

	sz := len(slice1)
	for i := 0; i < sz; i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
