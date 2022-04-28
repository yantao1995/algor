package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2248 lang=golang
 *
 * [2248] 多个数组求交集
 */

// @lc code=start
func intersection(nums [][]int) []int {
	m := map[int]int{}
	for k := range nums {
		for kk := range nums[k] {
			m[nums[k][kk]]++
		}
	}
	result := []int{}
	for k := range m {
		if m[k] == len(nums) {
			result = append(result, k)
		}
	}
	sort.Ints(result)
	return result
}

// @lc code=end

/*
	m记录每个数出现的次数
	如果次数等于 nums 的长度，就说明都出现过
	然后排序输出
*/
