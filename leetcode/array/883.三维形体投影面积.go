package leetcode

/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] 三维形体投影面积
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	top, x, y := 0, 0, 0
	mi, mj := map[int]int{}, map[int]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				top++
				if grid[i][j] > mj[i] {
					mj[i] = grid[i][j]
				}
				if grid[i][j] > mi[j] {
					mi[j] = grid[i][j]
				}
			}
		}
	}
	for i := range mi {
		x += mi[i]
	}
	for i := range mj {
		y += mj[i]
	}
	return top + x + y
}

// @lc code=end

/*
	top 记录顶部面积  x 记录x面积   y记录y面积
	x取x看过去的最高边，相同 i 的最高v
	y取y看过去的最高边，相同 j 的最高v
	顶部只需要值大于0就能看得到面积

*/
