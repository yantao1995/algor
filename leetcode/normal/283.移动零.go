package leetcode

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		temp := i
		for temp > 0 {
			if nums[temp-1] == 0 {
				nums[temp], nums[temp-1] = nums[temp-1], nums[temp]
				temp--
			} else {
				goto comehere
			}
		}
	comehere:
	}
}

// @lc code=end
