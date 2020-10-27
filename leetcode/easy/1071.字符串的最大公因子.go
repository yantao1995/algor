package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	shortStr, longStr := str1, str2
	if len(str2) < len(str1) {
		shortStr, longStr = str2, str1
	}
	containSlice := []string{}
	for i := len(shortStr); i > 0; i-- {
		if len(shortStr)%i == 0 {
			tempStr := shortStr[:i]
			for j := i; j+i <= len(shortStr); j += i {
				if tempStr != shortStr[j:j+i] {
					goto come
				}
			}
			containSlice = append(containSlice, tempStr)
		}
	come:
	}
	for _, sli := range containSlice {
		longSli := strings.Split(longStr, sli)
		for _, v := range longSli {
			if v != "" {
				goto comeHere
			}
		}
		return sli
	comeHere:
	}
	return ""
}

// @lc code=end
