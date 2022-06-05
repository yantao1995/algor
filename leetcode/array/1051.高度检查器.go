package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1051 lang=golang
 *
 * [1051] 高度检查器
 */

// @lc code=start
func heightChecker(heights []int) int {
	heightsCopy := make([]int, len(heights))
	copy(heightsCopy, heights)
	sort.Ints(heightsCopy)
	result := 0
	for k := range heights {
		if heights[k] != heightsCopy[k] {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	普通 解法
    就是看每位与当前排序后数据的比较，不一样就是位置错了
*/
