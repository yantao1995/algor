package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1838 lang=golang
 *
 * [1838] 最高频元素的频数
 */

// @lc code=start
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	maxCount := 1
	currenctTotal := 0
	l, r := 0, 1
	for ; r < len(nums); r++ {
		currenctTotal += (nums[r] - nums[r-1]) * (r - l)
		for currenctTotal > k {
			currenctTotal -= nums[r] - nums[l]
			l++
		}
		if r-l+1 > maxCount {
			maxCount = r - l + 1
		}
	}
	return maxCount
}

// @lc code=end
