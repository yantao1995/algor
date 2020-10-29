package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	for k := range A {
		A[k] *= A[k]
	}
	sort.Ints(A)
	return A
}

// @lc code=end
