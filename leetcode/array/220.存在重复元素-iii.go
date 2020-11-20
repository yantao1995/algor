package leetcode

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] >= nums[j] && nums[i]-nums[j] <= t {
				return true
			}
			if nums[i] <= nums[j] && nums[j]-nums[i] <= t {
				return true
			}
		}
	}
	return false
}

// @lc code=end
