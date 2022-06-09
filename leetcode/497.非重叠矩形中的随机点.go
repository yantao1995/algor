package leetcode

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=497 lang=golang
 *
 * [497] 非重叠矩形中的随机点
 */

// @lc code=start
type Solution struct {
	rects [][]int
	r     *rand.Rand
}

func Constructor(rects [][]int) Solution {
	return Solution{
		rects: rects,
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Pick() []int {
	rtIndex := this.r.Intn(len(this.rects))
	return []int{
		this.rects[rtIndex][0] + this.r.Intn(this.rects[rtIndex][2]-this.rects[rtIndex][0]),
		this.rects[rtIndex][1] + this.r.Intn(this.rects[rtIndex][3]-this.rects[rtIndex][1]),
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
// @lc code=end
