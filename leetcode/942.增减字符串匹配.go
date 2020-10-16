package leetcode

/*
 * @lc app=leetcode.cn id=942 lang=golang
 *
 * [942] 增减字符串匹配
 */

// @lc code=start
func diStringMatch(S string) []int {
	result := make([]int, len(S)+1)
	index := 0
	for k := range S {
		if S[k] == 'I' {
			result[k] = index
			index++
		}
	}
	result[len(result)-1] = index
	index++
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == 'D' {
			result[i] = index
			index++
		}
	}
	return result
}

// @lc code=end
