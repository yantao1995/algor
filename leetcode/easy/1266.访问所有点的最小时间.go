package easy

/*
 * @lc app=leetcode.cn id=1266 lang=golang
 *
 * [1266] 访问所有点的最小时间
 */

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) int {
	result := [][]int{{points[0][0], points[0][1]}}
	for _, v := range points {
		for {
			last := len(result) - 1
			x, y := result[last][0], result[last][1]
			if v[0] > x {
				x += 1
			} else if v[0] < x {
				x -= 1
			}
			if v[1] > y {
				y += 1
			} else if v[1] < y {
				y -= 1
			}
			result = append(result, []int{x, y})
			if x == v[0] && y == v[1] {
				break
			}
		}
	}
	return len(result) - 2 //第一个多计算了一次 -1 ，两点之间距离为1，所以再 -1.
}

// @lc code=end
