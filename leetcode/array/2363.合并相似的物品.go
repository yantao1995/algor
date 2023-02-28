package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2363 lang=golang
 *
 * [2363] 合并相似的物品
 */

// @lc code=start
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, item := range append(items1, items2...) {
		m[item[0]] += item[1]
	}
	result := make([][]int, len(m))
	index := 0
	for k, v := range m {
		result[index] = []int{k, v}
		index++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

// @lc code=end

/*
	map 记录重量之和。然后result记录value和重量，升序即可
*/
