package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	sli := strings.Split(s, "")
	m := map[string]int{}
	for k := range sli {
		if _, ok := m[sli[k]]; !ok {
			m[sli[k]] = k
		} else {
			m[sli[k]] = len(sli) + 1
		}
	}
	min := len(sli) + 1
	for k := range m {
		if m[k] < min {
			min = m[k]
		}
	}
	if min == len(sli)+1 {
		return -1
	}
	return min
}

// @lc code=end
