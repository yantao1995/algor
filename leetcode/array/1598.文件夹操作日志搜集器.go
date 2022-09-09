package leetcode

/*
 * @lc app=leetcode.cn id=1598 lang=golang
 *
 * [1598] 文件夹操作日志搜集器
 */

// @lc code=start
func minOperations(logs []string) int {
	result := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if result > 0 {
				result--
			}
		} else if logs[i] != "./" {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	遇到文件夹就计数加1
	遇到 ../ 就减，最少减0
*/
