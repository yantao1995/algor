package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := min(height[l], height[r]) * (r - l)
	for l+1 < r {
		if height[l] > height[r] {
			for r > 1 {
				r--
				maxArea = max(maxArea, min(height[l], height[r])*(r-l))
				if height[r] > height[l] {
					break
				}
			}
		} else {
			for l+1 < len(height)-1 {
				l++
				maxArea = max(maxArea, min(height[l], height[r])*(r-l))
				if height[l] > height[r] {
					break
				}
			}
		}
	}
	return maxArea
}

// @lc code=end

/*
优化，双指针，向内层循环


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
	根接雨水不同的是，这个把边想象成无限薄，即忽略柱子的影响
	两次 双重循环，第一次以左边界为最低的边界，第二次以右边界为最低边界
	第一次，i找左边界，j找右边界来计算当前的最大盛水量
	第二次，倒过来
*/
