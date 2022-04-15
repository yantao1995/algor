package leetcode

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp1 := make([]int, len(nums)-1) //包含0，不包含末尾
	dp2 := make([]int, len(nums)-1) //不包含0，包含末尾
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(nums); i++ {
		if i >= 1 { //dp1
			if i == 1 {
				dp1[i-1] = nums[i]
			} else if i == 2 {
				dp1[i-1] = max(nums[i], dp1[i-2])
			} else {
				dp1[i-1] = max(nums[i]+dp1[i-3], dp1[i-2])
			}
		}
		if i < len(nums)-1 { //dp2
			if i == 0 {
				dp2[i] = nums[i]
			} else if i == 1 {
				dp2[i] = max(dp2[i-1], nums[i])
			} else {
				dp2[i] = max(nums[i]+dp2[i-2], dp2[i-1])
			}
		}
	}
	return max(dp1[len(dp1)-1], dp2[len(dp2)-1])
}

// @lc code=end

/*
	nums [1,2,3,4]
	dp1 记录索引 1到结尾
	dp2 记录索引 0到结尾-1
	这样就规避了首位相连，即 首和尾 分别取与不取
*/
