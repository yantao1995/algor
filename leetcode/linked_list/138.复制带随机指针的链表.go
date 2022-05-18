package leetcode

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	oriMap, nowMap := map[*Node]int{}, map[int]*Node{}
	for h, idx := head, 0; h != nil; h, idx = h.Next, idx+1 {
		oriMap[h] = idx
	}
	idx := 0
	for h := head; h != nil; h, idx = h.Next, idx+1 {
		node := &Node{
			Val:    h.Val,
			Next:   nil,
			Random: nil,
		}
		nowMap[idx] = node
	}
	for i, h := 0, head; i < idx; i, h = i+1, h.Next {
		nowMap[i].Next = nowMap[i+1]
		if ii, ok := oriMap[h.Random]; h.Random != nil && ok {
			nowMap[i].Random = nowMap[ii]
		}
	}
	return nowMap[0]
}

// @lc code=end

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
	两个map分别记录原始link的结点和索引，新link的结点和索引
	第一个for创建原始结点与索引关系
	第二个for创建新link
	第三个for根据原始结点的randmon指针指向的索引，关联当前link
*/
