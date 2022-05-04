package leetcode

/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	min, max := nums[0], nums[0]
	for k := range nums {
		if min > nums[k] {
			min = nums[k]
		}
		if max < nums[k] {
			max = nums[k]
		}
	}
	if max-min > 2*k {
		return max - min - 2*k
	}
	return 0
}

// @lc code=end

/*
找到最大最小值，因为两个数的计算是可以加k或者减k的，所以只要两个数的差值在2k内，都可以通过加减将差值转化成0，所以大于2k，最小差值就应该是 max-min-2*k
*/
