package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1636 lang=golang
 *
 * [1636] 按照频率将数组升序排序
 */

// @lc code=start
func frequencySort(nums []int) []int {
	count := map[int]int{}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		return count[nums[i]] < count[nums[j]] ||
			(count[nums[i]] == count[nums[j]] && nums[i] > nums[j])
	})
	return nums
}

// @lc code=end

/*
	count 记录每个数的频率，然后排序
*/
