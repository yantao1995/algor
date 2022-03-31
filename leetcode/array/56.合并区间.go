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

/*

	用count来标记重合区间，题目说0 <= start_i <= end_i <= 1e4，所以直接初始化line标记数组为 1e4 的长度。
	用 endPoint 来记录最大区间的右端点，获取结果的时候，不用只需要遍历到endPoint即可。
	标记：遍历所有区间，单区间左值如果在标记数组count上有值，则直接取该值覆盖区间。
	取值：遍历标记数组count,相同count值即为同一区间。
*/
