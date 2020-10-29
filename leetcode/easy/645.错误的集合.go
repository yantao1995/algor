package leetcode

/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	ht := map[int]int{}
	for k := range nums {
		ht[nums[k]]++
	}
	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if ht[i] > 1 {
			result = append(result, i)
		}
	}
	for i := 1; i <= len(nums); i++ {
		if ht[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end
