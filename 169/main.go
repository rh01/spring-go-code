package main

func majorityElementv1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	var m = make(map[int]int)
	n := len(nums)
	for _, v := range nums {
		m[v]++
		if m[v] > n / 2 {
			return v 
		}
	}
	return -1
}


func majorityElement(nums []int) int {
	count, ret := 0, 0
	for _, v := range nums {
		if count == 0 {
			ret = v
			count++
		} else {
			if v == ret {
				count++
			} else {
				count--
			}
		}
	}

	return ret
}

func main() {
	
}