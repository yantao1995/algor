package leetcode

/*
 * @lc app=leetcode.cn id=1822 lang=golang
 *
 * [1822] 数组元素积的符号
 */

// @lc code=start
func arraySign(nums []int) int {
	minusCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			minusCount++
		}
	}
	if minusCount%2 == 1 {
		return -1
	}
	return 1
}

// @lc code=end

/*
	记录正负数的个数，如果遇到0，直接返回0即可
	如果负数个数为单数，直接返回 -1
	否则返回0
*/
