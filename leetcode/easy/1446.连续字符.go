package leetcode

/*
 * @lc app=leetcode.cn id=1446 lang=golang
 *
 * [1446] 连续字符
 */

// @lc code=start
func maxPower(s string) int {
	lastIndex := 0
	maxCount := 0
	for k := range s {
		if s[k] != s[lastIndex] {
			if k-lastIndex > maxCount {
				maxCount = k - lastIndex
			}
			lastIndex = k
		}
	}
	if len(s)-lastIndex > maxCount {
		maxCount = len(s) - lastIndex
	}
	return maxCount
}

// @lc code=end
