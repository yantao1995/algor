package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	sli := strings.Split(s, " ")
	length := len(sli)
	for k := range sli {
		if sli[k] == "" {
			length--
		}
	}
	return length
}

// @lc code=end
