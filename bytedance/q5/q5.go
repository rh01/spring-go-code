package q5

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	i, j, k, c := 0, 0, l1+l2-2, 0 // 分别表示nums1和nums2的索引，以及k表示xxm,c表示进位carry
	arr := make([]int, l1+l2)      // 存放乘积
	out := make([]byte, 0)         // 存放结果
	for i, n1 := range []byte(num1) {
		for j, n2 := range []byte(num2) {
			arr[k-i-j] += int((n1 - '0') * (n2 - '0'))
		}
	}

	for i := 0; i < l1+l2; i++ {
		arr[i] += c
		c = arr[i] / 10
		arr[i] = arr[i] % 10
	}

	i = l1 + l2 - 1

	// 从头开始扫描，直到不为0
	for i >= 0 && arr[i] == 0 {
		i--
	}

	// 如果i此时小于0，则直接返回0即可
	if i < 0 {
		return "0"
	}

	// 下面开始计算返回的值
	for i >= 0 {
		out = append(out, byte(arr[i]+'0'))
		i--
	}

	return string(out)
}
