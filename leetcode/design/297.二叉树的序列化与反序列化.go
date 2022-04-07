package leetcode

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

// type Codec struct {
// 	list    *list.List
// 	buffer  bytes.Buffer
// 	tree    *TreeNode
// 	getNode func(val string) *TreeNode
// }

// func Constructor() Codec {
// 	return Codec{
// 		list:   list.New(),
// 		buffer: bytes.Buffer{},
// 		tree:   nil,
// 		getNode: func(val string) *TreeNode {
// 			if val == "null" {
// 				return nil
// 			}
// 			n0, _ := strconv.Atoi(val)
// 			return &TreeNode{
// 				Val:   n0,
// 				Left:  nil,
// 				Right: nil,
// 			}
// 		},
// 	}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	this.list.Init()
// 	this.buffer.Reset()
// 	this.list.PushBack(root)
// 	for this.list.Len() > 0 {
// 		element := this.list.Front()
// 		if this.buffer.Len() > 0 {
// 			this.buffer.WriteByte(',')
// 		}
// 		if tree, ok := element.Value.(*TreeNode); ok {
// 			if tree != nil {
// 				this.buffer.WriteString(strconv.Itoa(tree.Val))
// 				this.list.PushBack(tree.Left)
// 				this.list.PushBack(tree.Right)
// 			} else {
// 				this.buffer.WriteString("null")
// 			}
// 		}

// 		this.list.Remove(element)
// 	}
// 	return this.buffer.String()
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	this.tree = nil
// 	this.list.Init()
// 	if data == "" {
// 		return this.tree
// 	}
// 	nodes := strings.Split(data, ",")
// 	for i := 0; i < len(nodes); {
// 		if i == 0 {
// 			this.tree = this.getNode(nodes[i])
// 			this.list.PushBack(this.tree)
// 		}
// 		if element := this.list.Front(); element != nil {
// 			if node, ok := element.Value.(*TreeNode); ok && node != nil {
// 				i++
// 				if i < len(nodes) {
// 					node.Left = this.getNode(nodes[i])
// 					if node.Left != nil {
// 						this.list.PushBack(node.Left)
// 					}
// 				}
// 				i++
// 				if i < len(nodes) {
// 					node.Right = this.getNode(nodes[i])
// 					if node.Right != nil {
// 						this.list.PushBack(node.Right)
// 					}
// 				}
// 			} else {
// 				i++
// 			}
// 			this.list.Remove(element)
// 		} else {
// 			i++
// 		}
// 	}
// 	return this.tree
// }

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
