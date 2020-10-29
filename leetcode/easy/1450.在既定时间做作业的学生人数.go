package leetcode

/*
 * @lc app=leetcode.cn id=1450 lang=golang
 *
 * [1450] 在既定时间做作业的学生人数
 */

// @lc code=start
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for k := range startTime {
		if queryTime >= startTime[k] && queryTime <= endTime[k] {
			count++
		}
	}
	return count
}

// @lc code=end
