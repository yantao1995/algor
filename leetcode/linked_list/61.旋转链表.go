package leetcode

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	length := 0
	for h := head; h != nil; h = h.Next {
		length++
	}
	k %= length
	if k == 0 {
		return head
	}
	k = length - k
	if k == 0 {
		return head
	}
	length = 0
	var temp *ListNode
	for h := head; h != nil; h = h.Next {
		k--
		if k == 0 {
			temp = h
			break
		}
	}
	hd := temp.Next
	temp.Next = nil
	for h := hd; ; h = h.Next {
		if h.Next == nil {
			h.Next = head
			break
		}
	}
	return hd
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
