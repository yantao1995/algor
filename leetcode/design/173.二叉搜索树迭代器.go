package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */

//
type BSTIterator struct {
	Index  int
	Result []int
}

func Constructor(root *TreeNode) BSTIterator {
	result := []int{}
	lst := list.New()
	for lst.Len() > 0 || root != nil {
		if root != nil {
			lst.PushBack(root)
			root = root.Left
		} else {
			node := lst.Back()
			result = append(result, node.Value.(*TreeNode).Val)
			root = node.Value.(*TreeNode).Right
			lst.Remove(node)
		}
	}
	return BSTIterator{
		Index:  0,
		Result: result,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.Index >= len(this.Result) {
		return -1
	}
	this.Index++
	return this.Result[this.Index-1]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.Index < len(this.Result)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
