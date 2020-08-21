package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	maxCap, minHeight := 0, 0                                   //容量
	for interval := len(height) - 1; interval > 0; interval-- { //间隔
		for i := 0; i+interval < len(height); i++ {
			minHeight = height[i+interval]
			if height[i] < height[i+interval] {
				minHeight = height[i]
			}
			cap := minHeight * interval
			if cap > maxCap {
				maxCap = cap
			}
		}
	}
	return maxCap
}

// @lc code=end
