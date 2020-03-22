package utils

import (
	"bufio"
	"go/build"
	"os"
	"path/filepath"
	"strconv"
)

// GetArrayOfSize 获取指定大小的切片
func GetArrayOfSize(n int) []int {
	p, err := build.Default.Import("github.com/rh01/leetcode/00-algorithms/sort/utils", "", build.FindOnly)

	if err != nil {
		// handle error
	}

	fname := filepath.Join(p.Dir, "IntegerArray.txt")

	f, _ := os.Open(fname)

	defer f.Close()

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		s, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, s)
	}

	return numbers[0:n]
}
