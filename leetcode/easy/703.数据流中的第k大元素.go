package easy

/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 */

// @lc code=start
// type KthLargest struct {
// 	Pointer int
// 	Slice   []int
// 	Len     int
// }

// func Constructor(k int, nums []int) KthLargest {
// 	sort.Ints(nums)
// 	return KthLargest{
// 		Pointer: k,
// 		Slice:   nums,
// 		Len:     len(nums),
// 	}
// }

// func (this *KthLargest) Add(val int) int {
// 	this.Slice = append(this.Slice, val)
// 	this.Len++
// 	for i := this.Len - 1; i > 0 && this.Slice[i] < this.Slice[i-1]; i-- {
// 		this.Slice[i], this.Slice[i-1] = this.Slice[i-1], this.Slice[i]
// 	}
// 	return this.Slice[this.Len-this.Pointer]
// }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
