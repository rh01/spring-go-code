package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func appearTimeMax(arr []int, x int) int {
	var m = make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	var max = make([]int, 0)
	maxNum := 0
	for _, v := range m {
		if v > maxNum {
			maxNum = v
		}
	}

	for k, v := range m {
		if v == maxNum {
			max = append(max, k)
		}
	}

	for i := 0; i < len(max); i++ {
		for _, v := range arr {
			if v|x == max[i] {
				fmt.Println("in")
				return maxNum + 1
			}
		}
	}
	return maxNum
}

func main() {
	var m, n int
	// var i string

	fmt.Scanf("%d %d", &m, &n)
	//fmt.Scanln(&i)
	reader := bufio.NewReader(os.Stdin)
	// n, err := reader.ReadString('\n')
	// if n == "" || err != nil {
	// 	return
	// }
	// fmt.Println(n)

	bytes, _, _ := reader.ReadLine() //处理多行时，注意一下返回值的处理
	var nums = make([]int, m)
	i := 0
	fmt.Println(bytes)
	s := string(bytes)
	is := strings.Split(s, " ")
	for _, v := range is {
		nums[i], _ = strconv.Atoi(v)
		i++
	}

	fmt.Println(appearTimeMax(nums, n))

}
