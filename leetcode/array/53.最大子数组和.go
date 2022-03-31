package leetcode

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	max := -1 << 31
	current := 0
	for i := 0; i < len(nums); i++ {
		current += nums[i]
		if max < current {
			max = current
		}
		if current < 0 {
			current = 0
		}
	}
	return max
}

// @lc code=end

/*
	求子数组的最大值，i表示累加到第i个节点时的值
	如果当前i小于0，就抛弃当前累加值，从下一个位点开始计算
	负数  current 丢弃
	正数 就应该继续加
*/
