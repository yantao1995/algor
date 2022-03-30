package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	stack := []string{}
	result := ""
	last := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			if last != "" {
				stack = append(stack, last)
				last = ""
			}
		} else if s[i] == ']' {
			if last != "" {
				stack = append(stack, last)
			}
			last = "" //临时变量
			//每匹配成功一个[]就出栈
			for tail := len(stack) - 1; tail >= 0; tail-- {
				if stack[tail][0] >= 'a' && stack[tail][0] <= 'z' {
					last = stack[tail] + last
				} else {
					temp := ""
					for count, _ := strconv.Atoi(stack[tail]); count > 0; count-- {
						temp += last
					}
					stack = append(stack[0:tail], temp)
					break
				}
			}
			last = ""
		} else if s[i] >= '0' && s[i] <= '9' { //前面 括号 或者 字母
			if last == "" || last[0] >= '0' && last[0] <= '9' {
				last += string(s[i])
			} else {
				stack = append(stack, last)
				last = string(s[i])
			}
		} else if s[i] >= 'a' && s[i] <= 'z' { //前面必定为括号或者空
			last += string(s[i])
			if i == len(s)-1 {
				stack = append(stack, last)
			}
		}
	}
	for k := range stack {
		result += stack[k]
	}
	return result
}

// @lc code=end
