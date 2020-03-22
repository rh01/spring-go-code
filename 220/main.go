package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := map[int]int{}
	for i, num := range nums {
		for j, v := range m {
			if abs(v-num) <= t && abs(i-j) <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}

func abs(i0 int) int {
	if i0 < 0 {
		return -i0
	}
	return i0
}

func main() {

}
