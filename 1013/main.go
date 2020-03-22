package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	if A == nil || len(A) < 3 {
		return false
	}

	var sum int
	for _, v := range A {
		sum += v
	}

	fmt.Println(sum)

	// 若不能整除则返回false
	if sum%3 != 0 {
		return false
	}

	// 记录每一部分的和
	partSum := sum / 3

	// 两个指针
	i, j := 0, len(A)-1

	lsum := 0
	for i < j-1 {
		lsum += A[i]
		if lsum < partSum || lsum > partSum {
			i++
		} else {
			break
		}
	}
	fmt.Println(lsum)

	rsum := 0
	for j > i+1 {
		rsum += A[j]
		if rsum < partSum || rsum > partSum {
			j--
		} else {
			break
		}
	}
	fmt.Println(rsum)

	return rsum == partSum && lsum == partSum
}

func main() {

	arr := []int{18, 12, -18, 18, -19, -1, 10, 10}
	fmt.Println(canThreePartsEqualSum(arr))
}
