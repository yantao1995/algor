package leetcode

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] > max {
			max = heights[i]
		}
		minLine := heights[i]
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < minLine {
				minLine = heights[j]
			}
			if (j-i+1)*minLine > max {
				max = (j - i + 1) * minLine
			}
		}
	}
	return max
}

// @lc code=end
