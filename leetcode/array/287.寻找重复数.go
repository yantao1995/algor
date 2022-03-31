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

/*
	直接用map记录当前数字出现的频次，如果出现过，即找到这个重复数
*/
