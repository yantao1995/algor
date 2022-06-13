package leetcode

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */ //

// @lc code=start
func twoSum(numbers []int, target int) []int {
	for i := len(numbers) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{j + 1, i + 1}
			}
		}
	}
	return nil
}

// @lc code=end
