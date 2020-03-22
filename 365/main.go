package main

func canMeasureWater(x int, y int, z int) bool {
	return z == 0 || (x + y >= z && z % gcd(x, y) == 0)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {

}
