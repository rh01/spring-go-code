package sort

import (
	"fmt"
	"testing"

	"github.com/rh01/leetcode/00-algorithms/sort/utils"
)

func TestInserSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	insertSort(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			fmt.Println("list")
			t.Error()
		}
	}
}

func TestShellSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	shellSort(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			fmt.Println("list")
			t.Error()
		}
	}
}




func bentchmarkInsertSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		insertSort(list)
	}
}


func bentchmarkSelectSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		selectSort(list)
	}
}


func bentchmarkShellSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		shellSort(list)
	}
}


func BenchmarkInsertSort100(b *testing.B)   { bentchmarkInsertSort(100, b) }
func BenchmarkInsertSort1000(b *testing.B)  { bentchmarkInsertSort(1000, b) }
func BenchmarkInsertSort10000(b *testing.B) { bentchmarkInsertSort(10000, b) }


func BenchmarkSelectSort100(b *testing.B)   { bentchmarkSelectSort(100, b) }
func BenchmarkSelectSort1000(b *testing.B)  { bentchmarkSelectSort(1000, b) }
func BenchmarkSelectSort10000(b *testing.B) { bentchmarkSelectSort(10000, b) }




func BenchmarkShellSort100(b *testing.B)   { bentchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)  { bentchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B) { bentchmarkShellSort(10000, b) }