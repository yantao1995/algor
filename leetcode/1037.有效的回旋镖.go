package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	if len(points) == 2 {
		return true
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	val := float64(points[1][0]-points[0][0]) / float64(points[1][1]-points[0][1])
	for i := 1; i < len(points); i++ {
		if float64(points[i][0]-points[0][0])/float64(points[i][1]-points[0][1]) != val {
			return false
		}
	}
	return true
}

// @lc code=end
