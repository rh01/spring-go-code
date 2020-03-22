package lcof

import "testing"

func TestExchange(t *testing.T) {
	nums := []int{1, 2, 3, 4,5}
	if res := exchange(nums); res != nil {
		t.Errorf("%v", res)
	}
}
