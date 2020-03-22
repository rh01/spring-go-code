package main

func minArray(numbers []int) int {
	l, r := 0, len(numbers)

	for l <= r {
		mid := l + (r -l) >> 1
		if numbers[mid] <= numbers[r] {
			r = mid
		} else {
			r = mid
		}
	}
	return numbers[l]
}

func main() {
	
}