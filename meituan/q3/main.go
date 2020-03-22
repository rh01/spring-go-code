package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution(n, k, l, r int) int {
	// 前缀和取模
	// if (l+r)*(r-l+1)/2 != k {
	// 	return -1
	// }

	var sumArr = make([]int, r-l+2)
	sumArr[0] = 0
	j := 1
	for i := l; i <= r; i++ {
		sumArr[j] = sumArr[j-1] + i
		j++
	}

	return int(math.Pow(float64(r-l+1), float64(n)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// n, err := reader.ReadString('\n')
	// if n == "" || err != nil {
	// 	return
	// }
	// fmt.Println(n)

	bytes, _, _ := reader.ReadLine() //处理多行时，注意一下返回值的处理
	s := string(bytes)
	sarr := strings.Split(s, " ")
	params := make([]int, 4)
	for i, v := range sarr {
		params[i], _ = strconv.Atoi(v)
	}
	fmt.Println(params)

	fmt.Println(solution(params[0], params[1], params[2], params[3]))
}
