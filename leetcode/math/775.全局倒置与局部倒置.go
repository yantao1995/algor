package leetcode

/*
 * @lc app=leetcode.cn id=775 lang=golang
 *
 * [775] 全局倒置与局部倒置
 */

// @lc code=start
func isIdealPermutation(nums []int) bool {
	absf := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	for i := 0; i < len(nums); i++ {
		if absf(nums[i], i) > 1 {
			return false
		}
	}
	return true
}

// @lc code=end

/*
	数学归纳
	假设全都是顺序，那么都为0
	如果索引0位置的和索引1位置的交换，那么，刚好都为1相等。
	如果索引0和索引和索引2位置的交换，那么，局部为1，然后全局为2。
	所以当前的值如果和索引位置偏差超过1，那么就不能满足
*/
