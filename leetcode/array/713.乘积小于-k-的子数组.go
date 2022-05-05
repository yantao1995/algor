package leetcode

/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于 K 的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	total := 0
	for i := 0; i < len(nums); i++ {
		total = 1
		for j := i; j < len(nums); j++ {
			total *= nums[j]
			if total >= k {
				break
			}
			count++
		}
	}
	return count
}

// @lc code=end

/*
	双指针，i为起始位置，j为末尾位置
	total 记录当前累积积，count记录数量
	每轮 i 都置total为1，方便计算 nums[i] 单个数的积
*/
