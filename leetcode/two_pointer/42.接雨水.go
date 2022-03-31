package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	total := 0
	for i := 0; i < len(height)-1; i++ {
		sub := 0
		if height[i] > 0 {
			flag, maxIndex, subMax, subTemp := false, i+1, 0, 0
			for j := i + 1; j < len(height); j++ {
				if j == i+1 && height[j] > height[i] {
					flag = true
					break
				}
				if height[j] < height[i] {
					if height[j] >= height[maxIndex] {
						maxIndex, subMax = j, subTemp
					}
					sub, subTemp = sub+height[j], subTemp+height[j]
				} else {
					total += (height[i])*(j-i-1) - sub
					i, flag = j-1, true
					break
				}
			}
			if !flag && height[maxIndex] > 0 && maxIndex != i {
				total += (height[maxIndex])*(maxIndex-i-1) - subMax
				i = maxIndex - 1
			}
		}
	}
	return total
}

// @lc code=end

/*
	双指针，
	i为左边，j为右边
	找边的时候，边必须大于0
	可能存在左边大，右边小的情况，所以遍历j的时候，同时记录两种值：
	a：j>i时，直接取当前j，然后得到雨水，同时左边界移动到j。
	b：未找到j>i时，此时判定左边大于右边，应该取i右边最大的一个值来当右边。
	以a为优先级最高来取值，未取到则取b的值
*/
