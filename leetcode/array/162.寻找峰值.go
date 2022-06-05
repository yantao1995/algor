package leetcode

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if len(nums) == 1 || nums[i] > nums[i+1] {
				return 0
			}
		} else if i == len(nums)-1 {
			if nums[i] > nums[i-1] {
				return i
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return 0
}

// @lc code=end

/*
	直接一遍扫描，然后左右比较
*/
