package easy

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */
//

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tempN := n
	pointers := []*ListNode{}
	n++ //记录前一个指针
	h := head
	for head != nil {
		if n != 0 {
			pointers = append(pointers, head)
			n--
		} else {
			temp := pointers[len(pointers)-1]
			next := head
			for i := len(pointers) - 1; i >= 0; i-- { //指针集体后移
				temp = pointers[i]
				pointers[i] = next
				next = temp
			}
		}
		head = head.Next
	}
	if tempN < 0 || n > 1 {
		return h
	} else {
		if n == 1 {
			return h.Next
		}
		if tempN > 1 {
			pointers[0].Next = pointers[2]
		} else {
			pointers[0].Next = nil
		}
	}
	return h
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
