package leetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 1
	currentValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > currentValue {
			nums[length], currentValue = nums[i], nums[i]
			length++
		}
	}
	return length
}

// @lc code=end
