package lcof

import (
	"reflect"
	"testing"
)

func TestPrintNumbers(t *testing.T) {
	n := 1
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if res := printNumbers(n); !reflect.DeepEqual(want, res) {
		t.Errorf("get %v", res)
	}
}

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，
// 则打印出 1、2、3 一直到最大的 3 位数 999。

// 示例 1:

// 输入: n = 1
// 输出: [1,2,3,4,5,6,7,8,9]
