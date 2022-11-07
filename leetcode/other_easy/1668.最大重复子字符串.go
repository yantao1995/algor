package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=1668 lang=golang
 *
 * [1668] 最大重复子字符串
 */

// @lc code=start
func maxRepeating(sequence string, word string) int {
	result := 0
	temp := word
	for strings.Contains(sequence, temp) {
		result++
		temp += word
	}
	return result
}

// @lc code=end
