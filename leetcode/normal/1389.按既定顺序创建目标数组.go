package leetcode

/*
 * @lc app=leetcode.cn id=1389 lang=golang
 *
 * [1389] 按既定顺序创建目标数组
 */

// @lc code=start
func createTargetArray(nums []int, index []int) []int {
	result := []int{}
	for k := range nums {
		result = append(result, nums[k])
		temp := len(result) - index[k] - 1
		l := len(result) - 1
		for temp > 0 { //向前冒泡
			result[l], result[l-1] = result[l-1], result[l]
			l--
			temp--
		}
	}
	return result
}

// @lc code=end
