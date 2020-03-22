package lcof

import "strings"

func replaceSpaceV1(s string) string {
	newS := strings.Replace(s, " ", "%20", -1)
	return newS
}

func replaceSpace(s string) string {
	var res string
	for _, v := range s {
		switch v {
		case ' ':
			res += "%20"
		default:
			res += string(v)
		}
	}
	return res
}
