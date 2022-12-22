package leetcode

import (
	"sort"
)

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
	memo := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, len(nums))
		for j := i + 1; j < len(nums); j++ {
			memo[i][j] = gcd(nums[i], nums[j])
		}
	}
	result := 0
	route := map[int]bool{}
	var backtrace func(index int, scores, indexs []int)
	backtrace = func(index int, scores, indexs []int) {
		if len(scores)*2 == len(nums) {
			sort.Ints(scores)
			total := 0
			for i := 0; i < len(scores); i++ {
				total += (i + 1) * scores[i]
			}
			if total > result {
				result = total
			}
			return
		}
		for ; index < len(nums); index++ {
			if !route[index] {
				route[index] = true
				break
			}
		}
		for j := index + 1; j < len(nums); j++ {
			if !route[j] {
				route[j] = true
				backtrace(index+1, append(scores, memo[index][j]), append(indexs, index, j))
				route[j] = false
			}
		}
		route[index] = false
	}
	backtrace(0, nil, nil)
	return result
}

// @lc code=end

/*


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
