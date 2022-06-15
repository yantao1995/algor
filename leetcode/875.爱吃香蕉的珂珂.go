package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] >= piles[j]
	})

	return 0
}

// @lc code=end
