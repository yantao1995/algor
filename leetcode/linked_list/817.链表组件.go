package leetcode

/*
 * @lc app=leetcode.cn id=817 lang=golang
 *
 * [817] 链表组件
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	m := map[int]bool{}
	for k := range nums {
		m[nums[k]] = true
	}
	result := 0
	for head != nil {
		if m[head.Val] {
			for head != nil && m[head.Val] {
				head = head.Next
			}
			result++
		}
		if head != nil {
			head = head.Next
		}
	}
	return result
}

// @lc code=end

/*
	题目的意思是说：链表上连续的并且存在于nums里的一段链表视为一个组件，问一个有几个组件

	主要是题目比较绕，直接根据题意写即可。
*/
