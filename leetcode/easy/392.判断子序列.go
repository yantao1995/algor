package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	sli1 := strings.Split(s, "")
	sli2 := strings.Split(t, "")
	i, j := 0, 0
	for i < len(sli1) && j < len(sli2) {
		if sli1[i] == sli2[j] {
			i++
		}
		j++
	}
	if i != len(sli1) {
		return false
	}
	return true

}

// @lc code=end
