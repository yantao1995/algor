package leetcode

/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	popIndex := 0
	for pushIndex := 0; pushIndex < len(pushed); pushIndex++ {
		stack = append(stack, pushed[pushIndex])
		for len(stack) > 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}

// @lc code=end

/*
	popIndex 记录出栈索引
	pushIndex 记录入栈索引
	stack 记录当前栈内数据

	直接模拟整个过程，先入栈，然后循环只要遇到可以出栈的就出栈
	最后判断栈是否为空即可

*/
