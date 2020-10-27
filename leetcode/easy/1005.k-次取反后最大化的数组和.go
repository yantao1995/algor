package easy

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
func largestSumAfterKNegations(A []int, K int) int {
	result := 0
	sort.Ints(A)
	for i := 0; A[i] <= 0 && K > 0; i++ {
		A[i] = 0 - A[i]
		K--
	}
	min := 2<<32 - 1
	for k := range A {
		if min > A[k] {
			min = A[k]
		}
		result += A[k]
	}
	if K%2 == 1 && K > 0 {
		result -= 2 * min
	}
	return result
}

// @lc code=end
