package leetcode

/*
 * @lc app=leetcode.cn id=1260 lang=golang
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	rows, columns := len(grid), len(grid[0])
	k %= rows * columns //顺
	/*
		1 2 3
		4 5 6
		7 8 9
	*/
	for k > 0 {
		count := 0
		targetI, targetJ := 0, 0
		lastTargetI, lastTargetJ := 0, 0
		for {
			targetJ += (rows*columns - 1)
			if targetJ >= columns {
				targetI += targetJ / columns
				targetI %= rows
				targetJ %= columns
			}
			grid[targetI][targetJ], grid[lastTargetI][lastTargetJ] = grid[lastTargetI][lastTargetJ], grid[targetI][targetJ]
			count++
			if count == rows*columns {
				break
			}
			lastTargetI, lastTargetJ = targetI, targetJ
		}
		grid[targetI][targetJ], grid[lastTargetI][lastTargetJ] = grid[lastTargetI][lastTargetJ], grid[targetI][targetJ]
		k--
	}
	return grid
}

// @lc code=end
