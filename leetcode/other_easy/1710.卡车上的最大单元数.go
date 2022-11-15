package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1710 lang=golang
 *
 * [1710] 卡车上的最大单元数
 */

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	result := 0
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		if truckSize < boxTypes[i][0] {
			return result + truckSize*boxTypes[i][1]
		}
		result += boxTypes[i][0] * boxTypes[i][1]
		truckSize -= boxTypes[i][0]
	}
	return result
}

// @lc code=end

/*
	装箱数量固定，优先装单元数量大的，先按单元数量从大到小排序
	然后依次装即可
*/
