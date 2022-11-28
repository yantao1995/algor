package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=813 lang=golang
 *
 * [813] 最大平均值和的分组
 */

// @lc code=start
func largestSumOfAverages(nums []int, k int) float64 {
	acc := make([]float64, len(nums))
	acc[0] = float64(nums[0])
	for i := 1; i < len(nums); i++ {
		acc[i] = float64(nums[i]) + acc[i-1]
	}
	dp := make([][]float64, len(acc)) //i个数 k组
	for i := range dp {
		dp[i] = make([]float64, k+1)
		dp[i][1] = acc[i] / float64(i+1) //初始化列
	}
	for i := 1; i <= k; i++ {
		dp[0][i] = float64(nums[0]) //初始化行
	}

	for i := 1; i < len(nums); i++ {
		for kk := 2; kk <= k; kk++ {
			for j := 0; j < i; j++ {
				dp[i][kk] = math.Max(dp[i][kk], dp[j][kk-1]+(acc[i]-acc[j])/float64(i-j))
			}
		}
	}

	return dp[len(dp)-1][k]
}

// @lc code=end

/*
	背包变形 动态规划

	dp[i][j]  i表示i个数，j表示共分成j组
	初始化第0行和第1列即可
	转移方程：dp[i][kk] = max(dp[i][kk],dp[j][kk-1]+(acc[i]-acc[j])/float64(i-j))
			当前  =  max(当前 , kk-1组 0-j  + (i-j的和) )
*/
