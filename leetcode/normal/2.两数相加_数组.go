package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //

//方法名相同
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var node1, node2 []int
// 	for l1 != nil || l2 != nil {
// 		if l1 != nil {
// 			node1 = append(node1, l1.Val)
// 			l1 = l1.Next
// 		} else {
// 			node1 = append(node1, 0)
// 		}
// 		if l2 != nil {
// 			node2 = append(node2, l2.Val)
// 			l2 = l2.Next
// 		} else {
// 			node2 = append(node2, 0)
// 		}
// 	}
// 	length := len(node2)
// 	if len(node1) > len(node2) {
// 		length = len(node1)
// 	}
// 	kflag := 0
// 	for i := 0; i < length; i++ {
// 		k := node1[i] + node2[i] + kflag
// 		node1[i] = (k % 10)
// 		if k >= 10 {
// 			kflag = 1
// 		} else {
// 			kflag = 0
// 		}
// 	}
// 	if kflag == 1 {
// 		node1 = append(node1, 1)
// 	}
// 	var head *ListNode = &ListNode{Val: node1[0]}
// 	p := head
// 	for k, v := range node1 {
// 		if k != 0 {
// 			ln := ListNode{
// 				Val: v,
// 			}
// 			p.Next = &ln
// 			p = p.Next
// 		}
// 	}
// 	return head
// }

// @lc code=end
