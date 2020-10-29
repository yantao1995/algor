package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1200 lang=golang
 *
 * [1200] 最小绝对差
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	diff := 2<<31 - 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < diff {
			diff = arr[i] - arr[i-1]
		}
	}
	result := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == diff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}

// @lc code=end
