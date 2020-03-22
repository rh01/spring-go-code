package main

func distributeCandies(candies int, num_people int) []int {
	m, n, p := candies, 1, 0
	var res = make([]int, num_people)
	for 0 != m {
		if m <= n {
			res[p] += m
			break
		}
		m -= n
		res[p] += n
		n++
		p = (p + 1) % (num_people)
	}

	return res
}

func main() {

}
