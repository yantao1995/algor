package leetcode

/*
 * @lc app=leetcode.cn id=2249 lang=golang
 *
 * [2249] 统计圆内格点数目
 */

// @lc code=start
func countLatticePoints(circles [][]int) int {
	count := 0
	type dist struct {
		i, j int
	}
	m := map[dist]bool{}
	for _, crc := range circles {
		for xs, xe := crc[0]-crc[2], crc[0]+crc[2]; xs <= xe; xs++ {
			for ys, ye := crc[1]-crc[2], crc[1]+crc[2]; ys <= ye; ys++ {
				if (crc[0]-xs)*(crc[0]-xs)+(crc[1]-ys)*(crc[1]-ys) <= crc[2]*crc[2] {
					if !m[dist{xs, ys}] {
						count++
						m[dist{xs, ys}] = true
					}
				}
			}
		}
	}
	return count
}

// @lc code=end

/*
	当前点与 x 轴垂直， 然后 与圆心形成一个三角形 x^2 + y^2 = z^2
	如果当前点 与 圆心 的 的 z^2 小于或者等于 r^2 ，那么就在圆内
	最后用 map 来对重合圆来去重
*/
