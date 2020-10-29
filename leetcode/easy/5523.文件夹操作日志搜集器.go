package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=5523 lang=golang
 *
 * [5523] 文件夹操作日志搜集器
 */

// @lc code=start
func minOperations(logs []string) int {
	stack := list.New()
	for k := range logs {
		if logs[k] == "../" {
			if stack.Len() > 0 {
				stack.Remove(stack.Back())
			}
		} else if len(logs[k]) == 2 && logs[k][0] == '.' {
			continue
		} else {
			stack.PushBack(logs[k])
		}
	}
	return stack.Len()
}

// @lc code=end
