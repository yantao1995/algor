package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] &&
				intervals[i][1] <= intervals[j][1])

	})
	getMixed := func(a, b []int) (bool, []int) { //是否有交集
		result := make([]int, 2)
		if a[0] > b[0] {
			a, b = b, a
		}
		if a[1] <= b[0] {
			return false, result
		}
		result[0], result[1] = b[0], a[1]
		if b[1] < a[1] {
			result[1] = b[1]
		}
		return true, result
	}
	count := 0
	flag := true
	for i := 1; i < len(intervals); {
		flag, intervals[0] = getMixed(intervals[0], intervals[i])
		if flag {
			count++
		} else {
			intervals[0] = intervals[i]
		}
		i++
	}
	return count
}

// @lc code=end

/*
	和用最少数量的箭引爆气球 思路一样
	每次贪当前的最小重合
	排序之后，求当前的重合
*/
