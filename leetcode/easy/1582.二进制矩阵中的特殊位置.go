package leetcode

/*
 * @lc app=leetcode.cn id=1582 lang=golang
 *
 * [1582] 二进制矩阵中的特殊位置
 */

// @lc code=start
func numSpecial(mat [][]int) int {
	htRows, htCols := map[int]int{}, map[int]int{}
	count := 0
	for k := range mat {
		for j := range mat[k] {
			if mat[k][j] == 1 {
				htRows[k]++
				htCols[j]++
			}
		}
	}
	for k := range mat {
		for j := range mat[k] {
			if mat[k][j] == 1 && htRows[k] == 1 && htCols[j] == 1 {
				count++
			}
		}
	}
	return count
}

// @lc code=end
