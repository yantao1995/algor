package leetcode

/*
 * @lc app=leetcode.cn id=1147 lang=golang
 *
 * [1147] 段式回文
 */

// @lc code=start
func longestDecomposition(text string) int {
	for i := 1; i <= len(text)/2; i++ {
		if text[:i] == text[len(text)-i:] {
			return 2 + longestDecomposition(text[i:len(text)-i])
		}
	}
	if len(text) == 0 {
		return 0
	}
	return 1
}

// @lc code=end

/*
	贪心算法，每次都贪两端最短的字符串，然后递归向下求解
*/
