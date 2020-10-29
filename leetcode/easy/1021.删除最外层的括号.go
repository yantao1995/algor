package leetcode

/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	leftCount, lastIndex := 0, 0
	result := ""
	for k := range S {
		if S[k] == '(' {
			leftCount++
		} else {
			leftCount--
		}
		if leftCount == 0 {
			result += S[lastIndex+1 : k]
			lastIndex = k + 1
		}
	}
	return result
}

// @lc code=end
