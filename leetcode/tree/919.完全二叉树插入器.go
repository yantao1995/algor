package leetcode

/*
 * @lc app=leetcode.cn id=919 lang=golang
 *
 * [919] 完全二叉树插入器
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
type CBTInserter struct {
	t      *TreeNode
	queue  []*TreeNode
	isLeft bool
}

// func Constructor(root *TreeNode) CBTInserter {
// 	queue := []*TreeNode{root}
// 	for i := 0; i < len(queue); i++ {
// 		if queue[i].Left != nil {
// 			queue = append(queue, queue[i].Left)
// 		}
// 		if queue[i].Right != nil {
// 			queue = append(queue, queue[i].Right)
// 		}
// 		if queue[i].Left != nil && queue[i].Right != nil {
// 			queue = append(queue[:i], queue[i+1:]...)
// 			i--
// 		}
// 	}

// 	return CBTInserter{root, queue, queue[0].Left == nil}
// }

func (this *CBTInserter) Insert(val int) int {
	node := this.queue[0]
	newNode := &TreeNode{val, nil, nil}
	this.queue = append(this.queue, newNode)
	if this.isLeft {
		node.Left = newNode
	} else {
		node.Right = newNode
		this.queue = this.queue[1:]
	}
	this.isLeft = !this.isLeft
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.t
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
// @lc code=end

/*
 t 存储当前树的头结点
 queue 队列存储当前层序遍历的结点
 isLeft 标识 当前需要出队列的结点应该添加左结点还是右结点

 思路： 添加的子节点依次入队列，
 	   需要出队列的结点直到左右两个结点都被添加后，才让当前结点出队列
*/
