package leetcode

/*
 * @lc app=leetcode.cn id=1431 lang=golang
 *
 * [1431] 拥有最多糖果的孩子
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for k := range candies {
		if candies[k] > max {
			max = candies[k]
		}
	}
	result := make([]bool, len(candies))
	for k := range candies {
		if candies[k]+extraCandies >= max {
			result[k] = true
		}
	}
	return result
}

// @lc code=end
