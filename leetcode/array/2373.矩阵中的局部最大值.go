package leetcode

/*
 * @lc app=leetcode.cn id=2373 lang=golang
 *
 * [2373] 矩阵中的局部最大值
 */

// @lc code=start
func largestLocal(grid [][]int) [][]int {
	maxs := func(a, b, c []int) int {
		k := -1
		for _, v := range a {
			if k < v {
				k = v
			}
		}
		for _, v := range b {
			if k < v {
				k = v
			}
		}
		for _, v := range c {
			if k < v {
				k = v
			}
		}
		return k
	}
	result := make([][]int, len(grid)-2)
	for i := 1; i < len(grid)-1; i++ {
		result[i-1] = make([]int, len(grid)-2)
		for j := 1; j < len(grid)-1; j++ {
			result[i-1][j-1] = maxs(grid[i-1][j-1:j+2],
				grid[i][j-1:j+2], grid[i+1][j-1:j+2])
		}
	}
	return result
}

// @lc code=end

/*
	直接模拟全部过程即可
*/
