package leetcode

/*
 * @lc app=leetcode.cn id=1750 lang=golang
 *
 * [1750] 删除字符串两端相同字符后的最短长度
 */

// @lc code=start
func minimumLength(s string) int {
	i, j := 0, len(s)-1
	for s[i] == s[j] && i < j {
		for i+1 < j && s[i+1] == s[i] {
			i++
		}
		for j-1 > i && s[j-1] == s[j] {
			j--
		}
		i++
		j--
	}
	return j - i + 1
}

// @lc code=end

/*
	双指针，首尾相同时，向中间逼近，直到i>=j
*/
