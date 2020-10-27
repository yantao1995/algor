package easy

/*
 * @lc app=leetcode.cn id=1422 lang=golang
 *
 * [1422] 分割字符串的最大得分
 */

// @lc code=start
func maxScore(s string) int {
	count, maxCount := 0, 0
	for i := 1; i < len(s); i++ {
		count = 0
		for k := range s {
			if s[k] == '0' && k < i {
				count++
			}
			if s[k] == '1' && k >= i {
				count++
			}
		}
		if maxCount < count {
			maxCount = count
		}
	}
	return maxCount
}

// @lc code=end
