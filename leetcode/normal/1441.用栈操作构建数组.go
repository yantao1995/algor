package leetcode

/*
 * @lc app=leetcode.cn id=1441 lang=golang
 *
 * [1441] 用栈操作构建数组
 */

// @lc code=start
func buildArray(target []int, n int) []string {
	result := []string{}
	index := 0
	for i := 1; i <= n; i++ {
		if index >= len(target) {
			break
		} else if target[index] > i {
			result = append(result, []string{"Push", "Pop"}...)
		} else {
			result = append(result, "Push")
			index++
		}
	}
	return result
}

// @lc code=end
