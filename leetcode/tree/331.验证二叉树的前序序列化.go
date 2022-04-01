package leetcode

import (
	"container/list"
	"strings"
)

/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	if len(nodes) < 1 {
		return false
	}
	stack := list.New()
	for i := 0; i < len(nodes); i++ {
		stack.PushBack(nodes[i])
		for stack.Len() > 2 && stack.Back().Value.(string) == "#" && stack.Back().Prev().Value.(string) == "#" {
			stack.Remove(stack.Back())
			stack.Remove(stack.Back())
			if stack.Back().Value == "#" {
				return false
			}
			stack.Back().Value = "#"
		}
	}
	return stack.Len() == 1 && stack.Back().Value.(string) == "#"
}

// @lc code=end

/*
	使用栈来模拟操作
	遍历结点，依次入栈
	每次判断当前栈顶两个 #，即为叶子结点，
	然后依次出栈并设置当前尾结点为#，方便判断父节点
*/
