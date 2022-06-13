package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		sli := strings.Split(s, s[0:i])
		for k := range sli {
			if sli[k] != "" {
				goto come
			}
		}
		return true
	come:
	}
	return false
}

// @lc code=end
