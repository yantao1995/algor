package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=1408 lang=golang
 *
 * [1408] 数组中的字符串匹配
 */

// @lc code=start
func stringMatching(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}

// @lc code=end
