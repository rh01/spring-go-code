package lcof

import "testing"

func TestHammingWeight(t *testing.T) {
	n := 31
	if res := hammingWeight(n); res != 2 {
		t.Errorf("actual is %v, but got %v\n", 2, res)
	}
}
