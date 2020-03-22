package lcof

import "testing"

func TestFib(t *testing.T) {
	t.Log(2 << 31)
	if ret := fib(45); ret != 134903163 {
		t.Errorf("want %v, but got %v\n", 1, ret)
	}
}
