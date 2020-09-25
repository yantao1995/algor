package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	ht := map[int]int{}
	for k := range nums {
		ht[nums[k]]++
	}
	keys := make([]int, len(ht))
	ki := 0
	for k := range ht {
		keys[ki] = k
		ki++
	}
	sort.Ints(keys)
	maxLength := 0
	for k := range keys {
		if k+1 < len(keys) {
			if keys[k+1]-keys[k] == 1 {
				if ht[keys[k+1]]+ht[keys[k]] > maxLength {
					maxLength = ht[keys[k+1]] + ht[keys[k]]
				}
			}
		}
	}
	return maxLength
}

// @lc code=end
