package leetcode

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(slice []int, index int)
	backtrack = func(slice []int, index int) {
		result = append(result, append([]int{}, slice...))
		for i := index; i < len(nums); i++ {
			backtrack(append(slice, nums[i]), i+1)
		}
	}
	backtrack([]int{}, 0)
	return result
}

// @lc code=end

/*
	回溯，因为要求数组中元素互不相同
	所以记录当前的扫描的index，
	然后向后遍历来进行剪枝重复元素
*/
