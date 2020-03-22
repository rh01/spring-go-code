package lcof

import "testing"

func TestMinArray(t *testing.T) {
	array := []int{3, 4, 5, 1, 2}
	if res := minArray(array); res != 1 {
		t.Fatal("not matched")
	}
}
