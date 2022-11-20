package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=891 lang=golang
 *
 * [891] 子序列宽度之和
 */

// @lc code=start
func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	mod := int(1e9 + 7)
	result1, result2 := 0, 0
	for i := 0; i < len(nums); i++ {
		result1 = (result1*2 + nums[i]) % mod
	}
	for i := len(nums) - 1; i >= 0; i-- {
		result2 = (result2*2 + nums[i]) % mod
	}
	return (result2 - result1 + mod) % mod
}

// @lc code=end

/*
	求每个数对子序列的贡献度，最小值和最大值分别计算两次贡献度
	数学解法参考 ：秦九韶多项式

*/
