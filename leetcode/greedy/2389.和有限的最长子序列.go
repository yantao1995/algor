package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2389 lang=golang
 *
 * [2389] 和有限的最长子序列
 */

// @lc code=start
func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	result := make([]int, len(queries))
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for k := range result {
		i := 0
		for ; i < len(nums); i++ {
			if nums[i] > queries[k] {
				break
			}
		}
		result[k] = i
	}
	return result
}

// @lc code=end

/*
	子序列可以删除元素，所以相当于可以选择任意元素，
	那么可以使用贪心，组成最长子序列的一定是选择剩余元素中最小的，贪最小的元素,
	也就是可以升序处理原数组，然后用前缀和求解即可
*/
