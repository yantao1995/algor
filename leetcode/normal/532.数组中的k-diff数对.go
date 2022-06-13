package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 {
				if nums[j] == nums[j-1] {
					continue
				}
			}
			if nums[j]-nums[i] == k {
				count++
			}
			if nums[j]-nums[i] > k {
				break
			}
		}
	}
	return count
}

// @lc code=end
