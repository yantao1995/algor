package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1337 lang=golang
 *
 * [1337] 方阵中战斗力最弱的 K 行
 */

// @lc code=start
func kWeakestRows(mat [][]int, k int) []int {
	countSlice := make([]int, len(mat))
	countTemp := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		count := 0
		for k := range mat[i] {
			if mat[i][k] != 1 {
				break
			}
			count++
		}
		countSlice[i] = count
	}
	result := []int{}
	copy(countTemp, countSlice)
	sort.Ints(countTemp)
	for i := 0; i < k; i++ {
		for k := range countSlice {
			if countSlice[k] == countTemp[i] {
				result = append(result, k)
				countSlice[k] = -1
				break
			}
		}
	}
	return result
}

// @lc code=end
