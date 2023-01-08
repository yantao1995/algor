package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=2185 lang=golang
 *
 * [2185] 统计包含给定前缀的字符串
 */

// @lc code=start
func prefixCount(words []string, pref string) int {
	result := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	直接取前缀即可
*/
