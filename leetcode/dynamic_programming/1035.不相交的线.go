package leetcode

/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
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
	for i := 0; i < len(nums2); i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
		}
		if i > 0 {
			dp[0][i] = max(dp[0][i], dp[0][i-1])
		}

	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
		if i > 0 {
			dp[i][0] = max(dp[i][0], dp[i-1][0])
		}
	}

	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(nums1)-1][len(nums2)-1]
}

// @lc code=end

/*
	dp[i][j] 表示nums1[i]到nums2[j]的最大连线数
	因为每个数字智能属于一条连线，所以 nums1[i] 到 nums2[j]时，
	当前值如果可以连线，那么必然是在 nums1[i-1] 到 nums2[j-1]的基础上+1
	否则，任取 nums1[i-1]到nums2[j]  和 nums1[i]到nums2[j-1] 的最大值
*/
