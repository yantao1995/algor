package leetcode

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	lists := []*ListNode{}
	for h := head; h != nil; h = h.Next {
		lists = append(lists, h)
	}
	for i, j := 0, len(lists)-1; i < j; i, j = i+1, j-1 {
		lists[i].Next = lists[j]
		lists[j].Next = lists[i+1]
		lists[i+1].Next = nil
	}
}

// @lc code=end

/*
	用lists记录当前的所有节点，然后用i和j依次串联每个节点。
	为了避免出现循环连接，每次讲最后一个节点的 Next 置为 nil
*/
