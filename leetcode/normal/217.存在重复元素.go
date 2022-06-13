package leetcode

/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	ht := map[int]bool{}
	for k := range nums {
		if ht[nums[k]] {
			return true
		}
		ht[nums[k]] = true
	}
	return false
}

// @lc code=end
