package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] 寻找右区间
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	intervalsCopy := make([][]int, len(intervals))
	copy(intervalsCopy, intervals)
	for k := range intervalsCopy {
		intervalsCopy[k] = append([]int(nil), intervals[k]...)
		intervalsCopy[k][1] = k
	}
	sort.Slice(intervalsCopy, func(i, j int) bool {
		return intervalsCopy[i][0] < intervalsCopy[j][0]
	})
	result := make([]int, len(intervals))
	for k := range result {
		result[k] = -1
		for kk := range intervalsCopy {
			if intervals[k][1] <= intervalsCopy[kk][0] {
				result[k] = intervalsCopy[kk][1]
				break
			}
		}
	}
	return result
}

// @lc code=end

/*
	copy一个数组，用第二位来记录每个数组的当前索引,然后进行排序
	然后遍历原数组，再从copy数组中找到符合条件的
*/
