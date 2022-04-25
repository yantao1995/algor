package leetcode

import (
	"container/heap"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start

type Mod struct {
	Value  int //值
	Weight int //排序权值
}

type Ints []*Mod

func (n Ints) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Ints) Len() int           { return len(n) }
func (n Ints) Less(i, j int) bool { return n[i].Weight < n[j].Weight }

func (n *Ints) Push(x interface{}) {
	*n = append(*n, x.(*Mod))
}

func (n *Ints) Top() interface{} {
	old := *n
	return old[0]
}

func (n *Ints) Pop() interface{} {
	old := *n
	rtn := old[len(old)-1]
	*n = old[:len(old)-1]
	return rtn
}

func topKFrequent(nums []int, k int) []int {
	ht := map[int]int{}
	for k := range nums {
		ht[nums[k]]++
	}
	hp := &Ints{}
	heap.Init(hp)
	for key, v := range ht {
		if hp.Len() < k {
			md := &Mod{
				Value:  key,
				Weight: v,
			}
			heap.Push(hp, md)
		} else {
			if hp.Top().(*Mod).Weight < v {
				heap.Pop(hp)
				md := &Mod{
					Value:  key,
					Weight: v,
				}
				heap.Push(hp, md)
			}
		}

	}
	result := []int{}
	for _, v := range *hp {
		result = append([]int{v.Value}, result...)
	}
	return result
}

// @lc code=end
