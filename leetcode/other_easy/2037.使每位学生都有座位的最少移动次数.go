package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2037 lang=golang
 *
 * [2037] 使每位学生都有座位的最少移动次数
 */

// @lc code=start
func minMovesToSeat(seats []int, students []int) int {
	absf := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	result := 0
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < len(seats); i++ {
		result += absf(seats[i], students[i])
	}
	return result
}

// @lc code=end

/*
	升序，然后每个人都坐到当前座位，求绝对值
*/
