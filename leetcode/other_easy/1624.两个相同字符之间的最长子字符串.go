package leetcode

/*
 * @lc app=leetcode.cn id=1624 lang=golang
 *
 * [1624] 两个相同字符之间的最长子字符串
 */

// @lc code=start
func maxLengthBetweenEqualCharacters(s string) int {
	result := -1
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i && j > result+i+1; j-- {
			if s[i] == s[j] {
				result = j - i - 1
			}
		}
	}
	return result
}

// @lc code=end

/*
	外层循环从前往后，内层循环从后往前，直接找最大的长度即可
*/
