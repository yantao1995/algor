package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	max := 0
	for i := 0; i <= max && i < len(nums); i++ {
		if nums[i]+i > max {
			max = nums[i] + i
		}
	}
	return max+1 >= len(nums)
}

// @lc code=end

/*
	max用于记录当前可以走到的最大索引下标，
	然后遍历到max，获取下一轮能走到的最大下标，
	因为存在超过索引下标的情况，且len()=索引+1，
	所以 只要max+1 >= len()即可
*/
