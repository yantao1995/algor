package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	sqrtIndex := make([]int, int(math.Sqrt(float64(n)))+1)
	for i := 1; i < len(sqrtIndex); i++ {
		sqrtIndex[i] = i * i
	}
	minCount := 1
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }

	return minCount
}

// @lc code=end
// 1 4 9 16 25 36 49
