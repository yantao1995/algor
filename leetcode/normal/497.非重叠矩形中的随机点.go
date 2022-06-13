package leetcode

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=497 lang=golang
 *
 * [497] 非重叠矩形中的随机点
 */

// @lc code=start
type Solution struct {
	rects       [][]int //当前矩阵的点数范围
	weight      []int   //当前矩阵点的个数
	totalWeight int
	r           *rand.Rand
}

// func Constructor(rects [][]int) Solution {
// 	s := Solution{
// 		rects:  rects,
// 		weight: make([]int, len(rects)),
// 		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
// 	}
// 	for k := range s.weight {
// 		s.weight[k] = (rects[k][3] - rects[k][1] + 1) * (rects[k][2] - rects[k][0] + 1)
// 		s.totalWeight += s.weight[k]
// 		s.weight[k] = s.totalWeight - s.weight[k]
// 	}
// 	return s
// }

func (this *Solution) Pick() []int {
	rtIndex := this.r.Intn(this.totalWeight)
	for k := range this.weight {
		if k == len(this.weight)-1 || (this.weight[k] <= rtIndex && this.weight[k+1] > rtIndex) {
			rtIndex = k
			break
		}
	}
	return []int{
		this.rects[rtIndex][0] + this.r.Intn(this.rects[rtIndex][2]-this.rects[rtIndex][0]+1),
		this.rects[rtIndex][1] + this.r.Intn(this.rects[rtIndex][3]-this.rects[rtIndex][1]+1),
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
// @lc code=end

/*
	weight 记录当前矩阵的点数累积起始值， 比如  0  4  8  ，方便遍历时找到当前计算的权重所处的位置
	totalWeight 记录所有矩阵的点数的累计值，用于计算随机数应该处于哪个矩阵

	思路： 根据矩阵的点数加权来计算随机的矩阵，然后再在矩阵内随机一个值
	易错： 矩阵的点数 应该是  矩阵的高度+1 * 矩阵的宽度+1 ，比如长宽为1的矩阵，点数为4
*/
