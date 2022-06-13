package leetcode

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	ht := map[int]int{} //值，索引
	for m := range nums {
		if v, ok := ht[nums[m]]; ok { //存在
			if m-v <= k {
				return true
			}
		}
		ht[nums[m]] = m
	}
	return false
}

// @lc code=end
