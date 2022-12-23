package leetcode

/*
 * @lc app=leetcode.cn id=2011 lang=golang
 *
 * [2011] 执行操作后的变量值
 */

// @lc code=start
func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, op := range operations {
		if op[1] == '+' {
			result++
		} else {
			result--
		}
	}
	return result
}

// @lc code=end

/*
	直接模拟整个操作即可
*/
