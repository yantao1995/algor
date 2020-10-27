package easy

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func hasCycle(head *ListNode) bool {
	lowPoint, fastPoint := head, head
	for fastPoint != nil {
		fastPoint = fastPoint.Next
		if fastPoint != nil { //2格
			fastPoint = fastPoint.Next
		} else {
			return false
		}
		lowPoint = lowPoint.Next
		if fastPoint == lowPoint {
			return true
		}
	}
	return false
}

// @lc code=end
