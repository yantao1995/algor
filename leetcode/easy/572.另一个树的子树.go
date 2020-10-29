package leetcode

/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一个树的子树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	q := Queue572{}
	q.LPush(s)
	for q.Size != 0 {
		node := q.RPop().(*TreeNode)
		if node.Val == t.Val {
			if isEquals(node, t) {
				return true
			}
		}
		if node.Left != nil {
			q.LPush(node.Left)
		}
		if node.Right != nil {
			q.LPush(node.Right)
		}
	}
	return false
}

func isEquals(s, t *TreeNode) bool {
	if t == nil && s == nil {
		return true
	}
	if t != nil && s != nil {
		if t.Val == s.Val {
			return isEquals(s.Left, t.Left) && isEquals(s.Right, t.Right)
		}
	}
	return false
}

type Queue572 struct {
	QueueSlice []interface{}
	Head       int //队首
	Tail       int //队尾
	Size       int
}

//左入
func (q *Queue572) LPush(str interface{}) *Queue572 {
	if q.Head == 0 {
		if q.Size == len(q.QueueSlice) {
			s := []interface{}{str}
			s = append(s, q.QueueSlice...)
			q.QueueSlice = s
		} else { //右移
			for i := 0; i < q.Size; i++ {
				q.QueueSlice[q.Tail-i+1] = q.QueueSlice[q.Tail-i]
			}
		}
		if q.Size != 0 {
			q.Tail++
		}
	} else {
		if q.Size != 0 {
			q.Head--
		}
	}
	q.QueueSlice[q.Head] = str
	q.Size++
	return q
}

//右出
func (q *Queue572) RPop() interface{} {
	var rtnVal interface{}
	if q.Size > 0 {
		//fmt.Printf("%v\t", q.QueueSlice[q.Tail])
		rtnVal = q.QueueSlice[q.Tail]
		q.QueueSlice[q.Tail] = nil
		if q.Head != q.Tail {
			q.Tail--
		}
		q.Size--
		return rtnVal
	}
	return nil
}

// @lc code=end
