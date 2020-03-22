package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= (num - 1)
	}
	return res
}

func main() {
	i := uint32(11)
	fmt.Println(hammingWeight(i))
}
