package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	dp := make([][]int, len(stones))
	max := 0
	for k := range stones {
		if max < stones[k] {
			max = stones[k]
		}
	}
	for k := range dp {
		dp[k] = make([]int, len(stones))
		for kk := range dp[k] {
			dp[k][kk] = max
			if k == kk {
				dp[k][kk] = stones[k]
			}
		}
	}
	min := func(val int, nums ...int) int {
		for k := range nums {
			if val < nums[k] {
				val = nums[k]
			}
		}
		return val
	}
	getVal := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	getV := func(a, b, c int) int {
		return min(getVal(a, getVal(b, c)), getVal(b, getVal(a, c)), getVal(c, getVal(a, b)))
	}

	for i := 1; i < len(stones); i++ {
		dp[i-1][i] = getVal(stones[i-1], stones[i])
	}
	for i := len(stones) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if j == 0 {
				dp[j][i] = min(dp[j][i], getV(stones[i], stones[j], dp[1][i-1]))
			} else if j == i-1 {
				dp[j][i] = min(dp[j][i], getV(stones[i], stones[j], dp[0][j-1]))
			} else {
				dp[j][i] = min(dp[j][i], getV(getVal(stones[i], stones[j]), dp[0][j-1], dp[j+1][i-1]))
			}
		}
	}
	fmt.Println(dp)
	return dp[0][len(stones)-1]
}

// @lc code=end

/*

 */
