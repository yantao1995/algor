package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	start, end := 0, len(temp)
	flag := true
	for k := range nums {
		if nums[k] != temp[k] {
			start = k
			flag = false
			break
		}
	}
	if flag {
		return 0
	}
	for i := len(temp) - 1; i >= 0; i-- {
		if nums[i] != temp[i] {
			end = i
			break
		}
	}
	return end - start + 1
}

// @lc code=end
