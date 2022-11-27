package leetcode

/*
 * @lc app=leetcode.cn id=1752 lang=golang
 *
 * [1752] 检查数组是否经排序和轮转得到
 */

// @lc code=start
func check(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
		}
	}
	return count == 0 || (count == 1 && nums[0] >= nums[len(nums)-1])
}

// @lc code=end

/*
	可以轮换0次，那就是不存在前一个大于后一个元素的情况
	轮换到n个位置，那就是只会出现一个前一个大于后一个的情况
*/
