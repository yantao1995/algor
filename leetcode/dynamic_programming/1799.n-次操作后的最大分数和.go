package leetcode

import "math/bits"

/*
 * @lc app=leetcode.cn id=1799 lang=golang
 *
 * [1799] N 次操作后的最大分数和
 */

// @lc code=start
func maxScore(nums []int) int {
	//最大公约数 辗转相除法
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	memo := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, len(nums))
		for j := i + 1; j < len(nums); j++ {
			memo[i][j] = gcd(nums[i], nums[j])
		}
	}
	dp := make([]int, 1<<len(nums))
	for k := 0; k < 1<<len(nums); k++ {
		cnt := bits.OnesCount(uint(k))
		if cnt%2 == 0 {
			for i := 0; i < len(nums); i++ {
				if k>>i&1 == 1 {
					for j := i + 1; j < len(nums); j++ {
						if k>>j&1 == 1 {
							dp[k] = max(dp[k], dp[k^(1<<i)^(1<<j)]+cnt/2*memo[i][j])
						}
					}
				}
			}
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

/*
动态规划，参考官方题解

回溯 超时
Time Limit Exceeded
44/76 cases passed (N/A)
func maxScore(nums []int) int {
	//最大公约数 辗转相除法
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	memo := make([][]int, len(nums))
	for k := range memo {
		memo[k] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			memo[i][j] = gcd(nums[i], nums[j])
		}
	}
	result := 0
	route := map[int]bool{}
	var backtrace func(deep, scores int)
	backtrace = func(deep, scores int) {
		if deep*2 > len(nums) {
			if scores > result {
				result = scores
			}
			return
		}
		for i := 0; i < len(nums); i++ {
			if !route[i] {
				route[i] = true
				for j := i + 1; j < len(nums); j++ {
					if !route[j] {
						route[j] = true
						backtrace(deep+1, scores+deep*memo[i][j])
						route[j] = false
					}
				}
				route[i] = false
			}
		}
	}
	backtrace(1, 0)
	return result
}

*/
