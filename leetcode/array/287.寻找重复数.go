package leetcode

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	index := 0
	mCount := map[int]int{}
	for {
		if mCount[index] > 0 {
			return index
		}
		mCount[index]++
		index = nums[index]
	}
}

// @lc code=end
