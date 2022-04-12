package leetcode

/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	var backtrace func(index, cn, ck int, nums []int)
	backtrace = func(index, cn, ck int, nums []int) {
		if ck == k {
			if cn == n {
				result = append(result, append([]int{}, nums...))
			}
			return
		}
		for ; index < 10; index++ {
			backtrace(index+1, cn+index, ck+1, append(nums, index))
		}

	}
	backtrace(1, 0, 0, make([]int, 0, k))
	return result
}

// @lc code=end

/*
	回溯
	因为每个数字只能用一次，所以需要记录当前索引值
	遇到个数相同的时候，进行return判断
*/
