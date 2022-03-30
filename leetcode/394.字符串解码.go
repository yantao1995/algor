package leetcode

import (
	"fmt"
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
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			if last != "" {
				stack = append(stack, last)
				last = ""
			}
			count++
		} else if s[i] == ']' {
			count--
			if last != "" {
				stack = append(stack, last)
			}
			last = ""
			if count == 0 {
				//出栈
				fmt.Println("last:", last, "stack", stack)
				temp := ""
				for tail := len(stack) - 1; tail >= 0; tail-- {
					if stack[tail][0] >= '0' && stack[tail][0] <= '9' {
						if tail+2 < len(stack) {
							if stack[tail+2][0] >= '0' && stack[tail+2][0] <= '9' {
								stack[tail+1] += stack[tail+3]
							}
						}
						char := stack[tail+1]
						for count, _ := strconv.Atoi(stack[tail]); count > 1; count-- {
							stack[tail+1] += char
						}
						if tail+2 < len(stack) {
							if stack[tail+2][0] >= 'a' && stack[tail+2][0] <= 'z' {
								stack[tail+1] += stack[tail+2]
							}
						}
						temp = stack[tail+1]
						fmt.Println(stack[tail], stack[tail+1])
					}
				}
				stack = stack[0:0]
				result += temp
			}
		} else if s[i] >= '0' && s[i] <= '9' { //前面 括号 或者 字母
			if last == "" || last[0] >= '0' && last[0] <= '9' {
				last += string(s[i])
			} else {
				if count > 0 {
					stack = append(stack, last)
				} else {
					result += last
				}
				last = string(s[i])
			}
		} else if s[i] >= 'a' && s[i] <= 'z' { //前面必定为括号或者空
			last += string(s[i])
			if i == len(s)-1 {
				result += last
			}
		}
	}
	return result
}

// @lc code=end
