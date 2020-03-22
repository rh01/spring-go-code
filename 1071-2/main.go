package main

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	// 先判断是否有解
	if str1+str2 != str2+str1 {
		return ""
	}

	// 如果有解，求解gcd(len(str1), len(str2))即可
	return str1[0:gcd(len(str1), len(str2))]
}

func main() {

}
