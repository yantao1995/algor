package leetcode

/**
	给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
	给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
	如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
	如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。

 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
*/

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x, Next: nil}
	if aNode == nil {
		node.Next = node
		aNode = node
	} else if aNode.Next == aNode {
		aNode.Next = node
		node.Next = aNode
	} else {
		for h := aNode; ; h = h.Next {
			if (x > h.Val && x < h.Next.Val) ||
				(h.Next == aNode) ||
				(h.Val > h.Next.Val && (x < h.Next.Val || x > h.Val)) {
				node.Next = h.Next
				h.Next = node
				break
			}
		}
	}
	return aNode
}

/*
   判断条件：
	1.当前结点为 nil ，直接返回node
	2.下一个结点为头结点时，直接创建新节点
	3.两个结点以上时，可插入的条件：
		a.如果x大于当前，小于next
		b.如果当前结点的下一个结点是头结点
		c.如果当前结点的值大于下一个结点，那么x如果小于下一个结点或者大于当前结点
*/
