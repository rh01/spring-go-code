package lcof

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	tData := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	if w1 := findNumberIn2DArray(tData, 5); w1 != true {
		t.Errorf("wanted %v, expect %v\n", w1, true)
	}

}
