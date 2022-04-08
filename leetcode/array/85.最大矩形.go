package leetcode

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	max := 0
	dp := make([][]int, len(matrix))
	for k := range dp {
		dp[k] = make([]int, len(matrix[k]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				if i > 0 && matrix[i-1][j] == '1' {
					dp[i][j] = dp[i-1][j] + 1
				}
			}
		}
	}
	//行高度
	left, right := 0, 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] > 0 {
				left, right = j-1, j+1
				for left >= 0 && dp[i][left] >= dp[i][j] {
					left--
				}
				for right < len(dp[i]) && dp[i][right] >= dp[i][j] {
					right++
				}
				if max < (right-left-1)*dp[i][j] {
					max = (right - left - 1) * dp[i][j]
				}
			}
		}
	}
	return max
}

// @lc code=end

/*
	直接暴力会超时
	将矩形想象成一个一个的柱子,例如:
	1 0 1 0 0
	1 0 1 1 1
	1 1 1 1 1
	1 0 0 1 0
	可以转化成  dp数组记录的就是当前柱子的高度
	1 0 1 0 0
	2 0 2 1 1
	3 1 3 2 2
	4 0 0 3 0
	即从上倒下，依次递增
	此时，如果在第三行来看的话，如果矩形有相同的宽的话，那么只需要判断
	第三行的值是否大于等于当前的索引值
	例如 dp[2][2] 这一行，dp[2][1]比3小，所以不取,dp[2][3]比3小，也不取，那么当前值就是3
	     dp[2][3] 这一行，dp[2][2]比3大，可以取，然后向左移动继续判断,dp[2][4]等于3，也可以取，然后向右移动继续判断...
	对每一行都进行如上的判断，然后就可以找到最大的矩形的值。
*/
