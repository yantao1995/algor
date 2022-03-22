package leetcode

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	total := 0
	for k := range nums {
		total += nums[k]
	}
	if total%2 == 1 || len(nums) < 2 {
		return false
	}
	half := total / 2
	flag := false
	var backtrack func(current, index int)
	backtrack = func(current, index int) {
		if flag {
			return
		}
		if current == half {
			flag = true
			return
		}
		for i := index; i < len(nums) && !flag; i++ {
			backtrack(current+nums[i], i+1)
		}
	}
	backtrack(0, 0)
	return flag
}

// @lc code=end
