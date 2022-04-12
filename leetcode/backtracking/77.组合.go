package leetcode

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := [][]int{}
	var backtrace func(cn, ck int, nums []int)
	backtrace = func(cn, ck int, nums []int) {
		if ck == k {
			result = append(result, append([]int{}, nums...))
			return
		}
		for ; cn <= n; cn++ {
			backtrace(cn+1, ck+1, append(nums, cn))
		}
	}
	backtrace(1, 0, make([]int, 0, k))
	return result
}

// @lc code=end

/*
	回溯
	cn表示当前值,ck表示当前个数，然后向后挨个加
*/
