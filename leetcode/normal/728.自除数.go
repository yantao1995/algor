package leetcode

/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for ; left <= right; left++ {
		num := left
		for num > 0 {
			temp := num % 10
			if temp == 0 || left%temp != 0 {
				goto Skip
			}
			num /= 10
		}
		result = append(result, left)
	Skip:
	}
	return result
}

// @lc code=end
