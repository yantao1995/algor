package leetcode

/*
 * @lc app=leetcode.cn id=1785 lang=golang
 *
 * [1785] 构成特定和需要添加的最少元素
 */

// @lc code=start
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for k := range nums {
		sum += nums[k]
	}
	result := 0
	sub := goal - sum
	if sub < 0 {
		sub = -sub
	}
	if sub > 0 {
		result = sub / limit
		if sub%limit > 0 {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	求和，然后判断差了多少，然后对limit做除，可以得到需要添加的个数
	然后取余判断是否还需要再添加一个
*/
