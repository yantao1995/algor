package easy

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	ht2 := map[int]int{}
	for k := range arr2 {
		ht2[arr2[k]] = k + 1
	}
	result, notIn := []int{}, []int{}
	for k := range arr1 {
		if ht2[arr1[k]] != 0 {
			result = append(result, arr1[k])
		} else {
			notIn = append(notIn, arr1[k])
		}
	}
	sort.Ints(notIn)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if ht2[result[j]] < ht2[result[i]] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	result = append(result, notIn...)
	return result
}

// @lc code=end
