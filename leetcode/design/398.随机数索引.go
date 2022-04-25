package leetcode

/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

// @lc code=start
type Solution struct {
	m map[int][]int
	f func(n int) int
}

// func Constructor(nums []int) Solution {
// 	m := map[int][]int{}
// 	for k := range nums {
// 		if _, ok := m[nums[k]]; !ok {
// 			m[nums[k]] = []int{}
// 		}
// 		m[nums[k]] = append(m[nums[k]], k)
// 	}
// 	return Solution{
// 		m: m,
// 		f: func(n int) int {
// 			return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
// 		},
// 	}
// }

func (this *Solution) Pick(target int) int {
	return this.m[target][this.f(len(this.m[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

/*
	m : 记录当前元素的所有索引
	f : 获取数组长度的随机索引
	然后直接返回当前数据的索引数组的随机索引的值
*/
