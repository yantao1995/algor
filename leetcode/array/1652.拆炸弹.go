package leetcode

/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] 拆炸弹
 */

// @lc code=start
func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	for i := 0; i < len(result); i++ {
		if k > 0 {
			for j := 1; j <= k; j++ {
				result[i] += code[(i+j)%len(code)]
			}
		} else if k < 0 {
			tempi := i
			for j := 1; j <= -k; j++ {
				if tempi < j {
					tempi += len(code)
				}
				result[i] += code[tempi-j]
			}
		} else {
			break
		}
	}
	return result
}

// @lc code=end

/*
	直接累加即可
*/
