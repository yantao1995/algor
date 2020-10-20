package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */

// @lc code=start
func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i > 1; i-- {
		if A[i] < A[i-1]+A[i-2] { //周长连续
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

// @lc code=end
