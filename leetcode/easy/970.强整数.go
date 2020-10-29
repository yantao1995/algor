package leetcode

/*
 * @lc app=leetcode.cn id=970 lang=golang
 *
 * [970] 强整数
 */

// @lc code=start
func powerfulIntegers(x int, y int, bound int) []int {
	ht := map[int]bool{}
	result := []int{}
	for i := 1; i < bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			ht[i+j] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	for k := range ht {
		result = append(result, k)
	}
	return result
}

// @lc code=end
