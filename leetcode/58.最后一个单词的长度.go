package leetcode

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	lastMax, max := 0, 0
	for k := range s {
		if s[k] == ' ' {
			if max > lastMax {
				lastMax = max
			}
			max = 0
			continue
		}
		max++
	}
	if max > lastMax {
		return max
	}
	return lastMax
}

// @lc code=end
