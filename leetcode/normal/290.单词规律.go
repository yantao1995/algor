package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	patSli := strings.Split(pattern, "")
	sSli := strings.Split(s, " ")
	if len(patSli) != len(sSli) {
		return false
	}
	m := map[string]string{}
	for k := range patSli {
		if v, ok := m[patSli[k]]; !ok {
			for _, mv := range m {
				if mv == sSli[k] {
					return false
				}
			}
			m[patSli[k]] = sSli[k]
		} else {
			if v != sSli[k] {
				return false
			}
		}
	}
	return true
}

// @lc code=end
