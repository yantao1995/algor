package easy

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	for k > 0 {
		rotate1(nums)
		k--
	}

}

func rotate1(nums []int) { //1次
	nums[len(nums)-1], nums[0] = nums[0], nums[len(nums)-1]
	for i := len(nums) - 1; i > 1; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
}

// @lc code=end
