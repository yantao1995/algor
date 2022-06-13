package leetcode

/*
 * @lc app=leetcode.cn id=1413 lang=golang
 *
 * [1413] 逐步求和得到正数的最小值
 */

// @lc code=start
func minStartValue(nums []int) int {
	sum := 0
	minValue := 2<<31 - 1
	for k := range nums {
		sum += nums[k]
		if sum < minValue {
			minValue = sum
		}
	}
	if minValue < 0 {
		return 0 - minValue + 1
	}
	return 1
}

// @lc code=end
