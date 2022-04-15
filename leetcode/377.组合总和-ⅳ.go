package leetcode

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	for k := range dp {
		dp[k] = make([]int, target+1)
	}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < target+1; j++ {
			if i == 0 {
				if j%nums[i] == 0 {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = dp[i-1][j]
				if j == nums[i] {
					dp[i][j]++
				} else if j > nums[i] {
					dp[i][j] += 2 * dp[i][j-nums[i]]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1][target]
}

// @lc code=end
