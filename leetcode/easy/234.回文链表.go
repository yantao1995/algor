package leetcode

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func isPalindrome(head *ListNode) bool {
	sli := []int{}
	for head != nil {
		sli = append(sli, head.Val)
		head = head.Next
	}
	for i := 0; len(sli)-i > i; i++ {
		if sli[i] != sli[len(sli)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
