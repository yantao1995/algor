package leetcode

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	val := 0
	for i := 0; i < len(s); i++ {
		val *= 26
		val += int(s[i] - 64)
	}
	return val
}

// @lc code=end
