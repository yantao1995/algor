package leetcode

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */

//
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	nexts := []*ListNode{}
	val := 99999
	Index := 0
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			if lists[i].Val < val {
				Index = i
				val = lists[i].Val
			}
		}
		next := lists[i]
		nexts = append(nexts, next)
	}
	h := nexts[Index]
	if h == nil {
		return h
	}
	nexts[Index] = nexts[Index].Next
	head := h
	for {
		val = 99999
		Index = -99999
		for i := 0; i < len(nexts); i++ {
			if nexts[i] != nil {
				if nexts[i].Val < val {
					Index = i
					val = nexts[i].Val
				}
			}
		}
		if Index >= 0 {
			h.Next = nexts[Index]
			nexts[Index] = nexts[Index].Next
			h = h.Next
		} else {
			h = nil
			break
		}
	}
	return head
}

// @lc code=end
