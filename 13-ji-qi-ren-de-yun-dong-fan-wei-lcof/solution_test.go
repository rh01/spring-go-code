package lcof

import "testing"

func TestMovingCount(t *testing.T) {
	var m = 1
	var n = 2
	var k = 1
	if target := movingCount(m, n, k); target != 2 {
		t.Errorf("actual is %v, but got %v\n", 2, target)
	}
}
