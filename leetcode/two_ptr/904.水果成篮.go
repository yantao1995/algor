package leetcode

/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] 水果成篮
 */

// @lc code=start
func totalFruit(fruits []int) int {
	result := 0
	m := map[int]int{}
	for l, r := 0, 0; l+result < len(fruits); r++ {
		m[fruits[r]]++
		for len(m) > 2 {
			m[fruits[l]]--
			if m[fruits[l]] == 0 {
				delete(m, fruits[l])
			}
			l++
		}
		if r-l+1 > result {
			result = r - l + 1
		}
	}
	return result
}

// @lc code=end

/*
	题目读半天，种类0 1 2，不是种类个数。

	滑动窗口，使用m记录窗口内元素的个数，保证种类不大于2
*/
