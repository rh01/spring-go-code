package lcof

import "testing"

// 判断在一个矩阵中是否存在一条包含某字符串所有字符的路径
func TestExist(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
	}

	if ret := exist(board, "ABCE"); !ret {
		t.Fail()
	}
}
