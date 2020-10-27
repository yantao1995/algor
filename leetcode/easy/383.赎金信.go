package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	rs := strings.Split(ransomNote, "")
	ms := strings.Split(magazine, "")
	mm := map[string]int{}
	for k := range ms {
		mm[ms[k]]++
	}
	for k := range rs {
		if v, ok := mm[rs[k]]; ok {
			if v > 0 {
				mm[rs[k]]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
