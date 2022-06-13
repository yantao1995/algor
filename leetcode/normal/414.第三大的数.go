package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	m := map[int]bool{}
	int64Min := -1 << 63
	stack := []int{int64Min, int64Min, int64Min} // 小->大
	for k := range nums {
		if m[nums[k]] {
			continue
		}
		m[nums[k]] = true
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
	for k := range stack {
		if stack[k] == int64Min {
			return stack[2]
		}
	}
	return stack[0]
}

// @lc code=end
