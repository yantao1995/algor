package leetcode

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	for head.Val == val {
		if head.Next != nil {
			head = head.Next
		} else {
			return nil
		}
	}
	node := head
	prev := head
	for node != nil {
		if node.Val == val {
			if node.Next != nil {
				prev.Next = node.Next
				node = node.Next
			} else {
				prev.Next = nil
				return head
			}
		}
		prev = node
		node = node.Next
	}
	return head
}

// @lc code=end
