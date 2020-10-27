package easy

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */
//
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	next1, next2 := l1, l2
	h := l1
	if l1.Val > l2.Val {
		h = l2
		next2 = next2.Next
	} else {
		next1 = next1.Next
	}
	head := h
	for next1 != nil || next2 != nil {
		if next1 != nil && next2 != nil {
			if next1.Val < next2.Val {
				h.Next = next1
				next1 = next1.Next
			} else {
				h.Next = next2
				next2 = next2.Next
			}
		} else {
			if next1 == nil {
				h.Next = next2
				next2 = next2.Next
			} else {
				h.Next = next1
				next1 = next1.Next
			}
		}
		h = h.Next
	}
	h.Next = nil
	return head
}

// @lc code=end
