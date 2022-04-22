package leetcode

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([][]int, len(nums1))
	for k := range dp {
		dp[k] = make([]int, len(nums2))
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if dp[i][j] == 0 {
				for count := 0; i+count < len(nums1) && j+count < len(nums2) && nums1[i+count] == nums2[j+count]; count++ {
					dp[i+count][j+count] = count + 1
				}
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
		}
	}
	return dp[len(nums1)-1][len(nums2)-1]
}

// @lc code=end

/*
	dp[i][j] 记录 从 nums1从0到i 与 nums2从0到j 所能匹配的最大长度
	遇到当前值相等，且还未扫描过，即 dp[i][j] 为0 且 nums[i] == nums[j]
	就从当前值向右下角扫描，并对dp[i][j]赋值
*/
