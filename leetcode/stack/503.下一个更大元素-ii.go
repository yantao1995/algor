package leetcode

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0, len(nums)) //单调栈  栈顶到栈底递增
	result := make([]int, len(nums))
	find := make([]bool, len(nums))
	for j := 0; j < 2; j++ {
		for i := 0; i < len(nums); i++ {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if !find[i] {
					result[i] = -1
				}
				for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
					if !find[stack[len(stack)-1]] {
						result[stack[len(stack)-1]] = nums[i]
						find[stack[len(stack)-1]] = true
					}
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, i)
			}
		}
	}
	return result
}

// @lc code=end

/*
	使用单调栈，栈顶到栈底递增
	例：stack [3,2,1]的顺序
	单调栈记录了元素的索引
	遇到比栈顶索引元素大的时候就出栈，出栈索引元素的下一个更大元素就是即将入栈的元素
*/
