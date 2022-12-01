package leetcode

/*
 * @lc app=leetcode.cn id=1779 lang=golang
 *
 * [1779] 找到最近的有相同 X 或 Y 坐标的点
 */

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	absf := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	minDistance := 0
	index := -1
	for i := 0; i < len(points); i++ {
		if x == points[i][0] {
			if dis := absf(y, points[i][1]); index == -1 || dis < minDistance {
				index = i
				minDistance = dis
			}
		} else if y == points[i][1] {
			if dis := absf(x, points[i][0]); index == -1 || dis < minDistance {
				index = i
				minDistance = dis
			}
		}
	}
	return index
}

// @lc code=end

/*
	直接遍历找到最小的一个即可
*/
