package leetcode

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
func partition(head *ListNode, x int) *ListNode {
	left, right := []int{}, []int{}
	for h := head; h != nil; h = h.Next {
		if h.Val < x {
			left = append(left, h.Val)
		} else {
			right = append(right, h.Val)
		}
	}
	index1, index2 := 0, 0
	for h := head; h != nil; h = h.Next {
		if index1 < len(left) {
			h.Val = left[index1]
			index1++
		} else {
			h.Val = right[index2]
			index2++
		}
	}
	return head
}

// @lc code=end
