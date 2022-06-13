package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1331 lang=golang
 *
 * [1331] 数组序号转换
 */

// @lc code=start
func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	ht := map[int]int{}
	rank := 1
	ht[temp[0]] = rank
	for i := 1; i < len(temp); i++ {
		if temp[i] != temp[i-1] {
			rank++
		}
		ht[temp[i]] = rank
	}
	for k := range arr {
		arr[k] = ht[arr[k]]
	}
	return arr
}

// @lc code=end
