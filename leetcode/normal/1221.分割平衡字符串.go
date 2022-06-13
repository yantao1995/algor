package leetcode

/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

// @lc code=start
func balancedStringSplit(s string) int {
	count, countL, countR := 0, 0, 0
	for k := range s {
		if s[k] == 'L' {
			countL++
		} else {
			countR++
		}
		if countL == countR {
			count++
			countL, countR = 0, 0
		}
	}
	return count
}

// @lc code=end
