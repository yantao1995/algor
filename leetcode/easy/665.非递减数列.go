package easy

/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	needCahnge := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i > 0 {
				if i+2 < len(nums) {
					if (nums[i-1] > nums[i+2] || nums[i] > nums[i+2]) && nums[i-1] > nums[i+1] {
						return false
					}
				}
			}
			needCahnge++
		}
	}
	if needCahnge < 2 {
		return true
	}
	return false
}

// @lc code=end
