package leetcode

/*
 * @lc app=leetcode.cn id=1827 lang=golang
 *
 * [1827] 最少操作使数组递增
 */

// @lc code=start
/* func minOperations(nums []int) int {
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			result += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return result
} */

// @lc code=end

/*
	直接让当前元素刚好大于前一个元素，然后记录差值即可
*/
