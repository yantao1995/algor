package leetcode

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
// type NumArray struct {
// 	Array []int
// 	Sum   []int
// }

// func Constructor(nums []int) NumArray {
// 	total := 0
// 	sum := make([]int, len(nums))
// 	for i := range nums {
// 		total += nums[i]
// 		sum[i] = total
// 	}
// 	return NumArray{
// 		Array: nums,
// 		Sum:   sum,
// 	}
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	if i == 0 {
// 		return this.Sum[j]
// 	} else {
// 		return this.Sum[j] - this.Sum[i-1]
// 	}
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
