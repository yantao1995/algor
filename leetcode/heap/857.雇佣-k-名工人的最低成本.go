package leetcode

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode.cn id=857 lang=golang
 *
 * [857] 雇佣 K 名工人的最低成本
 */

// @lc code=start
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	minTotal, currentTotal, hp, costs := float64(1<<31), 0, &IntHeap857{}, make([]int, len(quality))
	for i := 0; i < len(quality); i++ {
		costs[i] = i
	}
	sort.Slice(costs, func(i, j int) bool { return wage[costs[i]]*quality[costs[j]] < wage[costs[j]]*quality[costs[i]] })
	for _, cost := range costs {
		currentTotal += quality[cost]
		heap.Push(hp, quality[cost])
		if hp.Len() > k {
			currentTotal -= heap.Pop(hp).(int)
		}
		if hp.Len() == k {
			if minTotal > float64(currentTotal)*float64(wage[cost])/float64(quality[cost]) {
				minTotal = float64(currentTotal) * float64(wage[cost]) / float64(quality[cost])
			}
		}
	}
	return minTotal
}

type IntHeap857 []int

func (h IntHeap857) Len() int           { return len(h) }
func (h IntHeap857) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap857) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap857) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap857) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// @lc code=end

/*
	参考官方题解


暴力超时： (44/46)
	按质量降序，单价降序，然后依次从单价索引开始取，
	每次都选能单价比当前单价低的最小 k-1 个质量最小的数，能保证达到最低期望工资
	并且都满足要求

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	for i := 0; i < len(quality); i++ { //质量降序
		for j := i + 1; j < len(quality); j++ {
			if quality[i] < quality[j] {
				wage[i], wage[j] = wage[j], wage[i]
				quality[i], quality[j] = quality[j], quality[i]
			}
		}
	}
	cost := make([]float64, len(quality))
	costSeq := make([]int, len(cost)) //单价降序 存索引
	for i := 0; i < len(cost); i++ {
		cost[i] = float64(wage[i]) / float64(quality[i])
		costSeq[i] = i
	}
	sort.Slice(costSeq, func(i, j int) bool {
		return cost[costSeq[i]] > cost[costSeq[j]]
	})
	minTotal := cost[costSeq[0]] * float64(quality[0]*k)
	currentTotal := 0.0
	count := 1
	distinctMap := map[int]bool{}
	for i := 0; i+k <= len(costSeq); i++ {
		currentTotal = float64(wage[costSeq[i]])
		count = 1
		for j := len(quality) - 1; j >= 0 && count < k; j-- {
			if !distinctMap[j] && j != costSeq[i] {
				count++
				currentTotal += cost[costSeq[i]] * float64(quality[j])
			}
		}
		if count == k && minTotal > currentTotal {
			minTotal = currentTotal
		}
		distinctMap[costSeq[i]] = true
	}
	return minTotal
}

*/
