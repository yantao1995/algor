package leetcode

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	tempSlice := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] < len(tempSlice) {
			tempSlice[nums[i]] = nums[i]
		}
	}
	for i := 1; i < len(tempSlice); i++ {
		if tempSlice[i] == 0 {
			return i
		}
	}
	return len(tempSlice)
}

// @lc code=end
