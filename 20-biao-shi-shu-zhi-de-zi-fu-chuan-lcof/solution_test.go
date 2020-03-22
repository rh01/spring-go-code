package lcof

import "testing"

func TestIsNumber(t *testing.T) {
	if res := isNumber("5e2"); !res {
		t.Errorf("got is %v", res)
	}
}

// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
// 例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"及"-1E-16"都表示数值，
// 但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
