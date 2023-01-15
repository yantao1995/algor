package leetcode

/*
 * @lc app=leetcode.cn id=2293 lang=golang
 *
 * [2293] 极大极小游戏
 */

// @lc code=start
func minMaxGame(nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for len(nums) > 1 {
		n := len(nums)
		temp := make([]int, n/2)
		for i := 0; i < n/2; i++ {
			if i&1 == 0 {
				temp[i] = min(nums[2*i], nums[2*i+1])
			} else {
				temp[i] = max(nums[2*i], nums[2*i+1])
			}
		}
		nums = temp
	}
	return nums[0]
}

// @lc code=end

/*
	直接模拟整个过程即可
*/
