package easy

/*
 * @lc app=leetcode.cn id=1103 lang=golang
 *
 * [1103] 分糖果 II
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	index := 0
	increase := 1
	result := make([]int, num_people)
	for candies > 0 {
		if candies >= increase {
			result[index] += increase
			candies -= increase
		} else {
			result[index] += candies
			candies = 0
		}
		index++
		increase++
		if index == num_people {
			index = 0
		}
	}
	return result
}

// @lc code=end
