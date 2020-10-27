package easy

/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	lastMax, current := 0, 0
	for k := range nums {
		if nums[k] == 1 {
			current++
		} else {
			if current > lastMax {
				lastMax = current
			}
			current = 0
		}
	}
	if current > lastMax {
		return current
	}
	return lastMax
}

// @lc code=end
