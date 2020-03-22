package lcof

import (
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	mock := []int{2, 3, 1, 0, 2, 3}
	expect := findRepeatNumber(mock)
	t.Logf("expect: %v\n", expect)
	wanted := 2
	if expect != wanted {
		t.Errorf("want %v, but expect %v\n", expect, wanted)
	}
}
