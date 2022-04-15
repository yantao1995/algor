package leetcode

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	memo := make([]bool, len(nums))
	for i := 0; i < len(nums)-1; {
		memo[i] = true
		max := nums[i] + i
		if max >= len(nums)-1 {
			break
		}
		tempJ := nums[i] + i
		for j := 0; j < len(nums) && j <= i+nums[i]; j++ {
			if max < nums[j]+j {
				max = nums[j] + j
				tempJ = j
			}
		}
		i = tempJ
		if tempJ < len(nums)-1 {
			memo[tempJ] = true
		}
	}
	count := 0
	for k := range memo {
		if memo[k] {
			count++
		}
	}
	return count
}

// @lc code=end

/*
	贪心:每次贪当前能达到的最大位置
	[2,3,1,1,4]
	比如当前 2,那么可以 走到 3  1 ，此时，就看3和1，哪个能走的更远
	然后就选3，然后到终点了，然后退出
*/
