package leetcode

/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	q := Queue637{}
	datas := []float64{}
	target, current := 1, 0
	tempData := []int{}
	tempTarget := 0
	q.LPush(root)
	for q.Size != 0 {
		node := q.RPop().(*TreeNode)
		//fmt.Printf("%v\t", node.Val)
		//fmt.Println("target:", target, "current:", current, "tempTarget", tempTarget)
		if node.Left != nil {
			q.LPush(node.Left)
		} else {
			tempTarget++
		}
		if node.Right != nil {
			q.LPush(node.Right)
		} else {
			tempTarget++
		}
		if current <= target {
			current++
			tempData = append(tempData, node.Val)
			if current == target {
				total := 0
				for k := range tempData {
					total += tempData[k]
				}
				datas = append(datas, float64(total)/float64(len(tempData)))
				tempData = []int{}
				current = 0
				target = target*2 - tempTarget
				tempTarget = 0
			}
		}
	}
	if current < target && current != 0 {
		total := 0
		for k := range tempData {
			total += tempData[k]
		}
		datas = append(datas, float64(total)/float64(len(tempData)))
	}

	return datas
}

type Queue637 struct {
	QueueSlice []interface{}
	Head       int //队首
	Tail       int //队尾
	Size       int
}

//左入
func (q *Queue637) LPush(str interface{}) *Queue637 {
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
func (q *Queue637) RPop() interface{} {
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
