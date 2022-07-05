package leetcode

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	mins := func(a int, b ...int) int {
		for k := range b {
			if b[k] < a {
				a = b[k]
			}
		}
		return a
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i-1]) {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += mins(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}
	return mins(triangle[len(triangle)-1][0], triangle[len(triangle)-1][1:]...)
}

// @lc code=end

/*
	动态规划：使用原triangle数组作为dp数组
	遍历当前行，使用上一行的可选的最小值作为累加
*/
