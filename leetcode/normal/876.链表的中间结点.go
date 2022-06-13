package leetcode

/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	h := head
	for h != nil {
		count++
		h = h.Next
	}
	count = count / 2
	for count > 0 {
		head = head.Next
		count--
	}
	return head
}

// @lc code=end
