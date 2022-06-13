package leetcode

/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 */

// @lc code=start
func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		for dist := 0; dist < len(A[i])/2; dist++ {
			A[i][dist], A[i][len(A[i])-1-dist] =
				A[i][len(A[i])-1-dist], A[i][dist]
		}
	}
	for i := range A {
		for j := range A[i] {
			if A[i][j] == 1 {
				A[i][j] = 0
			} else {
				A[i][j] = 1
			}
		}
	}
	return A
}

// @lc code=end
