package leetcode

/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	ht := map[byte]int{}
	word := "balloon"
	for k := range text {
		ht[text[k]]++
	}
	count := 0
	for {
		for k := range word {
			if ht[word[k]] == 0 {
				return count
			}
			ht[word[k]]--
		}
		count++
	}
}

// @lc code=end
