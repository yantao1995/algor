package leetcode

import (
	"bytes"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	queue   []*TreeNode
	getNode func(val string) *TreeNode
	bf      *bytes.Buffer
}

// func Constructor() Codec {
// 	return Codec{
// 		queue: []*TreeNode{},
// 		getNode: func(val string) *TreeNode {
// 			if val == "nil" {
// 				return nil
// 			}
// 			n0, _ := strconv.Atoi(val)
// 			return &TreeNode{
// 				Val:   n0,
// 				Left:  nil,
// 				Right: nil,
// 			}
// 		},
// 		bf: &bytes.Buffer{},
// 	}
// }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.bf.Reset()
	this.queue = []*TreeNode{}
	this.queue = append(this.queue, root)
	for len(this.queue) > 0 {
		if this.bf.Len() > 0 {
			this.bf.WriteRune(',')
		}
		tail := this.queue[0]
		if tail != nil {
			this.queue = append(this.queue, []*TreeNode{tail.Left, tail.Right}...)
			this.bf.WriteString(strconv.Itoa(tail.Val))
		} else {
			this.bf.WriteString("nil")
		}
		this.queue = this.queue[1:]
	}
	return this.bf.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodeStrs := strings.Split(data, ",")
	this.queue = []*TreeNode{}
	var tree *TreeNode

	tree = this.getNode(nodeStrs[0])
	this.queue = append(this.queue, tree)

	for i := 0; len(this.queue) > 0 && i < len(nodeStrs); {
		front := this.queue[0]
		i++
		if i < len(nodeStrs) {
			front.Left = this.getNode(nodeStrs[i])
			if front.Left != nil {
				this.queue = append(this.queue, front.Left)
			}
		}
		i++
		if i < len(nodeStrs) {
			front.Right = this.getNode(nodeStrs[i])
			if front.Right != nil {
				this.queue = append(this.queue, front.Right)
			}
		}
		this.queue = this.queue[1:]
	}
	return tree
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end

/*
	和二叉树的序列化一样，使用了广度优先搜索
	只不过此处将二叉树的序列化所使用的队列从
	抽象的内建container/list，换成了切片实现
*/
