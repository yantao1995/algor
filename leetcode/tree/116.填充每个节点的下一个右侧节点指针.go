package leetcode

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root, nil}
	for i := 0; i < len(queue); i++ {
		if queue[i] != nil {
			queue[i].Next = queue[i+1]
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		} else if i != len(queue)-1 {
			queue = append(queue, nil)
		}
	}
	return queue[0]
}

// @lc code=end

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
	二叉树，层序遍历，然后依次指向后一个节点
	末尾节点用nil补充,如果当前nil不处于末尾，直接补充nil节点
*/
