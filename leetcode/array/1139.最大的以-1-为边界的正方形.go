package leetcode

/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 */

// @lc code=start
func largest1BorderedSquare(grid [][]int) int {
	isSquare := func(x, y, length int) bool {
		for i := 0; i <= length; i++ {
			if grid[x][y+i] == 0 ||
				grid[x+i][y] == 0 ||
				grid[x+length][y+i] == 0 ||
				grid[x+i][y+length] == 0 {
				return false
			}
		}
		return true
	}
	length := len(grid)
	if len(grid[0]) < length {
		length = len(grid[0])
	}
	for ; length > 0; length-- {
		for i := 0; i+length <= len(grid); i++ {
			for j := 0; j+length <= len(grid[i]); j++ {
				if isSquare(i, j, length-1) {
					return length * length
				}
			}
		}
	}
	return length * length
}

// @lc code=end

/*
	直接依次枚举所有正方形，然后对边进行判断
*/
