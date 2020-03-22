package lcof

import "testing"

func TestCuttingRope(t *testing.T) {
	n := 10
	if res := cuttingRope(n); res != 36 {
		t.Fatalf("actual is %v, but got %v", 1, res)
	}

}
