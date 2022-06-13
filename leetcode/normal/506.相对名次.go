package leetcode

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
func findRelativeRanks(nums []int) []string {
	// "Gold Medal", "Silver Medal", "Bronze Medal"
	temp := make([]int, len(nums))
	result := make([]string, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	ht := map[int]int{}
	for k := range temp {
		ht[temp[k]] = len(temp) - k
	}
	for k := range nums {
		if ht[nums[k]] == 1 {
			result[k] = "Gold Medal"
		} else if ht[nums[k]] == 2 {
			result[k] = "Silver Medal"
		} else if ht[nums[k]] == 3 {
			result[k] = "Bronze Medal"
		} else {
			result[k] = strconv.Itoa(ht[nums[k]])
		}
	}
	return result
}

// @lc code=end
