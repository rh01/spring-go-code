package lcof

import (
	"strings"
)

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	strings.Replace(s, "E", "e", -1)

	parts := strings.Split(s, "e")
	l := len(parts)
	if l == 1 {
		// 判断是否该部分有数字，没有数字是不合法
		// 没有e的情况下，判断是否是整数或者浮点数即可
		return haveDigit(parts[0]) && isFloatOrInteger(parts[0])
	} else if l == 2 {
		// 判断两部分是否含有数字部分，没有就不合法
		// 第一部分可以是整数或者浮点数，第二部分必须是整数
		return haveDigit(parts[0]) && haveDigit(parts[1]) && isFloatOrInteger(parts[0]) && isInteger(parts[1])
	} else {
		return false
	}
}

func haveDigit(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}
	return false
}

func getNum(s string) (int, bool) {
	l := len(s)
	if l == 0 {
		return 0, false
	}
	ans := 0
	for i := 0; i < l; i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		} else {
			ans = ans*10 + int(s[i]-'0')
		}
	}
	return ans, true
}

func isFloatOrInteger(s string) bool {
	parts := strings.Split(s, ".")
	l := len(parts)
	if l == 1 {
		//如果长度是1，没有小数点，则一定是整数
		return isInteger(parts[0])
	} else if l == 2 {
		// 如果长度是2，则第一部分必然为整数，为了消除边界条件，例如输入为".1"或者"1."，填充数字消除边界条件
		// 小数点右面的数字，不能带符号
		return isInteger(parts[0]+"0") && isIntegerWithoutSign(parts[1]+"0")
	} else {
		return false
	}
}

func isInteger(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	index := 0
	if s[index] == '-' || s[index] == '+' {
		index++
	}

	return isIntegerWithoutSign(s[index:])
}

func isIntegerWithoutSign(s string) bool {
	if _, ok := getNum(s); ok {
		return true
	} else {
		return false
	}
}

// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
// 例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"及"-1E-16"都表示数值，
// 但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
