package leetcode

/*
 * @lc app=leetcode.cn id=1313 lang=golang
 *
 * [1313] 解压缩编码列表
 */

// @lc code=start
func decompressRLElist(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i += 2 {
		for nums[i] > 0 {
			result = append(result, nums[i+1])
			nums[i]--
		}
	}
	return result
}

// @lc code=end
