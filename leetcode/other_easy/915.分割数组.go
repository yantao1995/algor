package leetcode

/*
 * @lc app=leetcode.cn id=915 lang=golang
 *
 * [915] 分割数组
 */

// @lc code=start
func partitionDisjoint(nums []int) int {
	lMax := 0
lab:
	for i := 0; i < len(nums); i++ {
		if nums[i] > lMax {
			lMax = nums[i]
		}
		for j := i + 1; j < len(nums); j++ {
			if lMax > nums[j] {
				continue lab
			}
		}
		return i + 1
	}
	return 0
}

// @lc code=end

/*
	暴力搜索，因为题目表示至少能找到一个满足的
	外层循环记录 left 的长度，lMax 记录 left内的最大值
	内层循环查找 right 内的值，只要全部都比 lMax大，就满足条件，
	否则进入下一轮搜索
*/
