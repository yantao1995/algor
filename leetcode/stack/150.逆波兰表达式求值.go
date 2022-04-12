package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for k := range tokens {
		switch tokens[k] {
		case "+":
			stack = append(stack[:len(stack)-2],
				stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2],
				stack[len(stack)-2]-stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2],
				stack[len(stack)-2]*stack[len(stack)-1])
		case "/":
			stack = append(stack[:len(stack)-2],
				stack[len(stack)-2]/stack[len(stack)-1])
		default:
			iInt, _ := strconv.Atoi(tokens[k])
			stack = append(stack, iInt)
		}
	}
	return stack[0]
}

// @lc code=end

/*
	数字就入栈，遇到操作符就取 len(stack)-2 和 len(stack)-1 的数字做运算
*/
