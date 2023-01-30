package leetcode

/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var l2h, l2t *ListNode = list2, list2
	for l2t.Next != nil {
		l2t = l2t.Next
	}
	var l1h, l1t *ListNode
	for i, node := 0, list1; i <= b; i++ {
		if i == a-1 {
			l1h = node
		}
		if i == b {
			l1t = node.Next
		}
		node = node.Next
	}
	if a == 0 {
		l2t.Next = l1t
		return l2h
	}
	l1h.Next = l2h
	l2t.Next = l1t
	return list1
}

// @lc code=end

/*
	先找到 list2 的头尾结点
	然后找 list1 的a前一个位置的结点和b位置的下一个结点
	再看a是否为0，如果为0，那么新链表的头应该是list2的头，否则为list1的头
*/
