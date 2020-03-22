package substring

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	target := "bab"
	if s2 := longestPalindrome(s); s2 != target {
		fmt.Println(s2)
		t.Error()
	}

}
