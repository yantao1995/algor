package leetcode

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	current, count := 0, 0
	for i := 0; i < len(nums); i++ {
		current = 0
		for j := i; j < len(nums); j++ {
			current += nums[j]
			if current == k {
				count++
			}
		}
	}
	return count
}

// @lc code=end
