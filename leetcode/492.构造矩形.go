package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	L := int(math.Sqrt(float64(area)))
	for L <= area {
		if area%L == 0 {
			if L >= area/L {
				return []int{L, area / L}
			}
		}
		L++
	}
	return []int{}
}

// @lc code=end
