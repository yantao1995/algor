package leetcode

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	source, dest []int
	rd           *rand.Rand
}

// func Constructor(nums []int) Solution {
// 	return Solution{
// 		source: nums,
// 		dest:   append([]int(nil), nums...),
// 		rd:     rand.New(rand.NewSource(int64(time.Now().Nanosecond()))),
// 	}
// }

func (this *Solution) Reset() []int {
	return this.source
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.dest); i++ {
		j := this.rd.Intn(len(this.dest))
		this.dest[i], this.dest[j] = this.dest[j], this.dest[i]
	}
	return this.dest
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

/*
	source 记录原始数据
	dest 记录每次打乱的顺序的数据
	rd 随机数生成器，每次随机生成数据的位置
*/
