package easy

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := 0, 0
	ha, hb := headA, headB
	for ha != nil {
		ha = ha.Next
		la++
	}
	for hb != nil {
		hb = hb.Next
		lb++
	}
	limit := la
	diff := lb - la
	ha, hb = headA, headB
	if la > lb {
		limit = lb
		diff = la - lb
		for diff > 0 {
			ha = ha.Next
			diff--
		}
	} else {
		for diff > 0 {
			hb = hb.Next
			diff--
		}
	}
	for limit > 0 {
		if ha == hb {
			return ha
		}
		ha = ha.Next
		hb = hb.Next
		limit--
	}
	return nil
}

// @lc code=end
