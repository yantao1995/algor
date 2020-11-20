package leetcode

/*
 * @lc app=leetcode.cn id=885 lang=golang
 *
 * [885] 螺旋矩阵 III
 */

// @lc code=start
func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	//1东  2南  3西  4北
	direct := 1
	length := 0
	result := [][]int{}
	last1, last2, last3, last4 := 0, 0, 0, 0
	length1, length2, length3, length4 := 1, 1, 2, 2
	if r0 < R && r0 >= 0 && c0 < C && c0 >= 0 {
		result = append(result, []int{r0, c0})
	}
	for len(result) < R*C {
		if direct == 1 {
			length1 += last1
			length = length1
			last1 = 2
		} else if direct == 2 {
			length2 += last2
			length = length2
			last2 = 2
		} else if direct == 3 {
			length3 += last3
			length = length3
			last3 = 2
		} else {
			length4 += last4
			length = length4
			last4 = 2
		}
		for temp := length; temp > 0; temp-- {
			if direct == 1 {
				c0++
			} else if direct == 2 {
				r0++
			} else if direct == 3 {
				c0--
			} else {
				r0--
			}
			if r0 < R && r0 >= 0 && c0 < C && c0 >= 0 {
				result = append(result, []int{r0, c0})
			}
		}
		direct++
		if direct == 5 {
			direct = 1
		}
	}
	return result
}

// @lc code=end
