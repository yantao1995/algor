package leetcode

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i-1] + 1
		}
	}
	return nums
}

// @lc code=end
