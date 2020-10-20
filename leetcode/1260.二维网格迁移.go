package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=1260 lang=golang
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	rows, columns := len(grid)-1, len(grid[0])-1
	fmt.Println("rows, columns", rows, columns)
	k %= rows * columns
	/*
		1 2 3
		4 5 6
		7 8 9
	*/
	temp := 0
	targetI, targetJ := 0, 0
	lastTargetI, lastTargetJ := 0, 0
	for {
		targetJ += k
		if targetJ >= columns {
			if targetJ == columns {
				targetI++
				targetI %= rows
				targetJ = 0
			} else {
				targetI += targetJ / columns
				targetI %= rows
				targetJ %= columns
			}
		}
		temp = grid[targetI][targetJ]
		grid[targetI][targetJ] = grid[lastTargetI][lastTargetJ]
		fmt.Print("targetI,targetJ ", targetI, targetJ)
		fmt.Print("\t lasttargetI,lasttargetJ ", lastTargetI, lastTargetJ)
		fmt.Print("\t temp ", temp)
		println()
		if targetI == 0 && targetJ == 0 {
			grid[targetI][targetJ] = temp
			return grid
		}
		lastTargetI, lastTargetJ = targetI, targetJ
	}
}

// @lc code=end
