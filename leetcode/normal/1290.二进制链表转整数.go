package leetcode

/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result *= 2
		result += head.Val
		head = head.Next
	}
	return result
}

// @lc code=end
