package leetcode

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}

// @lc code=end
