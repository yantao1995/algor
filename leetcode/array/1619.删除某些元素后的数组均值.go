package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1619 lang=golang
 *
 * [1619] 删除某些元素后的数组均值
 */

// @lc code=start
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	start, end := int(float64(len(arr))*0.05), int(float64(len(arr))*0.95)-1
	length := end - start + 1
	total := 0.0
	for ; start <= end; start++ {
		total += float64(arr[start])
	}
	return total / float64(length)
}

// @lc code=end

/*
	排序，然后累加
*/
