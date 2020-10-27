package easy

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
func lastStoneWeight(stones []int) int { //最大的两块互相消耗
	for len(stones) > 1 {
		sort.Ints(stones)
		stones = append(stones[:len(stones)-2], stones[len(stones)-1]-stones[len(stones)-2])
	}
	return stones[0]
}

// @lc code=end
