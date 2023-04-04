package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1637 lang=golang
 *
 * [1637] 两点之间不包含任何点的最宽垂直面积
 */

// @lc code=start
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	result := 0
	for i := 1; i < len(points); i++ {
		if points[i][0]-points[i-1][0] > result {
			result = points[i][0] - points[i-1][0]
		}
	}
	return result
}

// @lc code=end

/*
	排序后求x轴的距离即可
*/
