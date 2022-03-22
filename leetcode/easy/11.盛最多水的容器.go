package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	maxCap, cap := 0, 0
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			if height[j] >= height[i] {
				if cap = (j - i) * height[i]; cap > maxCap {
					maxCap = cap
				}
				break
			}
		}
	}
	for i := len(height) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if height[j] >= height[i] {
				if cap = (i - j) * height[i]; cap > maxCap {
					maxCap = cap
				}
				break
			}
		}
	}
	return maxCap
}

// @lc code=end
