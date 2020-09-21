package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	max := nums[0]
	if len(nums) < 3 {
		for k := range nums {
			if nums[k] > max {
				max = nums[k]
			}
		}
		return max
	}
	stack := []int{-9999, -9999, -9999} // 小->大
	for k := range nums {
		if nums[k] > stack[0] {
			stack[0] = nums[k]
			if stack[1] < stack[0] {
				stack[1], stack[0] = stack[0], stack[1]
			}
			if stack[2] < stack[1] {
				stack[2], stack[1] = stack[1], stack[2]
			}
			fmt.Println(stack)
		}
	}
	fmt.Println(stack)
	return stack[0]
}

// @lc code=end
