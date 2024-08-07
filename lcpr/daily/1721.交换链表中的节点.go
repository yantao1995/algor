package lcpr

/*
 * @lc app=leetcode.cn id=1721 lang=golang
 * @lcpr version=30204
 *
 * [1721] 交换链表中的节点
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	var prev, next *ListNode
	n := 0
	for h := head; h != nil; h = h.Next {
		n++
		if n == k {
			prev = h
		}
	}

	for h := head; h != nil; h = h.Next {
		if n == k {
			next = h
		}
		n--
	}
	prev.Val, next.Val = next.Val, prev.Val
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [7,9,6,6,7,8,3,0,9,5]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n2\n
// @lcpr case=end

*/

/*
	数总个数，然后记录前后指针，直接交换即可
*/
