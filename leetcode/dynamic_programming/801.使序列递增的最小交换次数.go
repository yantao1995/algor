package leetcode

/*
 * @lc app=leetcode.cn id=801 lang=golang
 *
 * [801] 使序列递增的最小交换次数
 */

// @lc code=start
func minSwap(nums1 []int, nums2 []int) int {
	dp := make([][2]int, len(nums1))
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[0][1] = 1
	for i := 1; i < len(nums1); i++ {
		dp[i][0] = len(nums1)
		dp[i][1] = len(nums1)
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			dp[i][0] = min(dp[i][0], dp[i-1][0])
			dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = min(dp[i][0], dp[i-1][1])
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
		}
	}
	return min(dp[len(nums1)-1][0], dp[len(nums1)-1][1])
}

// @lc code=end

/*
	参考官方题解
		因为满足题意，
			1.存在  nums1[i]>nums1[i−1] 且 nums2[i] > nums2[i - 1]nums2[i]>nums2[i−1] , 所以需要当前位置和前一个位置 同时交换或者不交换
			2.存在  nums1[i]>nums2[i−1] 且 nums2[i] > nums1[i - 1]nums2[i]>nums1[i−1] ，所以当前位置和前一位置只能其一发生交换

*/
