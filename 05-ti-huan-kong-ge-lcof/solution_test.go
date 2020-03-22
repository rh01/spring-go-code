package lcof

import "testing"

func TestReplaceSpace(t *testing.T) {
	ts := "We are happy."
	expect := "We%20are%20happy."
	if res := replaceSpace(ts); res != expect {
		t.Errorf("wanted %s, but get %s\n", res, expect)
	}

}
