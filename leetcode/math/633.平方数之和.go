package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	canSqrt := func(val int) bool {
		cSqrtFloat64 := math.Sqrt(float64(val))
		return cSqrtFloat64 == float64(int(cSqrtFloat64))
	}
	for left, right := 0, c; left*left <= c; left, right = left+1, c-(left+1)*(left+1) {
		if canSqrt(right) {
			return true
		}
	}
	return false
}

// @lc code=end

/*
	将C 分拆为  left平方 + right平方
	左边为平方数依次累加
*/
