package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	line := make([]uint8, 1e4)
	endPoint := 0
	result := [][]int{}
	count := uint8(1)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for k := range intervals {
		if intervals[k][1] > endPoint {
			endPoint = intervals[k][1]
		}
		if count+1 > 254 {
			count = 1
		} else {
			count++
		}
		if line[intervals[k][0]] >= 1 {
			count = line[intervals[k][0]]
		}
		for ; intervals[k][0] <= intervals[k][1]; intervals[k][0]++ {
			line[intervals[k][0]] = count
		}
	}
	for i := 0; i <= endPoint; i++ {
		if line[i] > 0 {
			for j := i; j <= endPoint; j++ {
				if line[j] != line[i] {
					result = append(result, []int{i, j - 1})
					i = j - 1
					break
				} else if j == endPoint {
					result = append(result, []int{i, j})
					i = j
				}
			}
		}
	}
	return result
}

// @lc code=end
