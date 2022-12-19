package leetcode

/*
 * @lc app=leetcode.cn id=1764 lang=golang
 *
 * [1764] 通过连接另一个数组的子数组得到一个数组
 */

// @lc code=start
func canChoose(groups [][]int, nums []int) bool {
	idx := 0
lab:
	for i := 0; idx < len(groups) && i <= len(nums)-len(groups[idx]); i++ {
		for j := 0; j < len(groups[idx]); j++ {
			if groups[idx][j] != nums[i+j] {
				continue lab
			}
		}
		i += len(groups[idx]) - 1
		idx++
	}
	return idx == len(groups)
}

// @lc code=end

/*
	像字符串匹配一样，挨个匹配
*/
