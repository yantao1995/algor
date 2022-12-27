package leetcode

/*
 * @lc app=leetcode.cn id=2027 lang=golang
 *
 * [2027] 转换字符串的最少操作次数
 */

// @lc code=start
func minimumMoves(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			result++
			i += 2
		}
	}
	return result
}

// @lc code=end

/*
	因为转换长度为3，
	所以可以直接遇到X就向后移动3个单位
	并且计数+1
*/
