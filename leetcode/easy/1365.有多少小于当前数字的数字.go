package easy

/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for k := range nums {
		for m := range nums {
			if nums[k] > nums[m] {
				result[k]++
			}
		}
	}
	return result
}

// @lc code=end
