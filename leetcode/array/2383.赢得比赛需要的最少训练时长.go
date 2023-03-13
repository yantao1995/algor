package leetcode

/*
 * @lc app=leetcode.cn id=2383 lang=golang
 *
 * [2383] 赢得比赛需要的最少训练时长
 */

// @lc code=start
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	minEnergy, minExperience := 0, 0
	min := func(a, b int) int {
		if b == 0 {
			b = -1
		}
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(energy); i++ {
		if initialEnergy <= energy[i] {
			minEnergy = min(minEnergy, initialEnergy-energy[i]-1)
		}
		if initialExperience <= experience[i] {
			minExperience = min(minExperience, initialExperience-experience[i]-1)
		}
		initialExperience += experience[i]
		initialEnergy -= energy[i]
	}
	return -(minEnergy + minExperience)
}

// @lc code=end

/*
	直接模拟整个过程即可
*/
