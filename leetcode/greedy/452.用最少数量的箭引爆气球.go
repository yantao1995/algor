package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=452 lang=golang
 *.
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	getMixed := func(a, b []int) (bool, []int) { //获取相交列表
		result := make([]int, 2)
		if a[0] > b[0] {
			a, b = b, a
		}
		if a[1] < b[0] {
			return false, result
		}
		result[0], result[1] = b[0], a[1]
		if b[1] < a[1] {
			result[1] = b[1]
		}
		return true, result
	}
	count := 1
	flag := true
	for i := 1; i < len(points); {
		flag, points[0] = getMixed(points[0], points[i])
		if !flag {
			points[0] = points[i]
			count++
		}
		i++
	}
	return count
}

// @lc code=end

/*
	按 points[i][0] 升序排列
	只要两两相交，那么就只需要一发箭
	然后如果相交，把相交的区域拿到，继续和下一个比较看看是不是两两相交
*/
