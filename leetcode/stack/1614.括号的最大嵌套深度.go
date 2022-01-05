package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=1614 lang=golang
 *
 * [1614] 括号的最大嵌套深度
 */

// @lc code=start
func maxDepth(s string) int {
	stack := list.New()
	max := 0
	for k := range s {
		if s[k] == '(' {
			stack.PushFront(s[k])
		} else if s[k] == ')' {
			if stack.Len() > 0 {
				if stack.Len() > max {
					max = stack.Len()
				}
				stack.Remove(stack.Front())
			}
		}
	}
	return max
}

// @lc code=end
