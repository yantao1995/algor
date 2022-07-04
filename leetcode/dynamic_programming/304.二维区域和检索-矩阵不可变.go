package leetcode

/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	dp [][][2]int //记录累加和  0行  1列
}

func Constructor(matrix [][]int) NumMatrix {
	dp := make([][][2]int, len(matrix))
	for k := range matrix {
		dp[k] = make([][2]int, len(matrix[k]))
		for kk := range dp[k] {
			dp[k][kk] = [2]int{}
			if k == 0 && kk == 0 {
				dp[k][kk][0] = matrix[k][kk]
				dp[k][kk][1] = matrix[k][kk]
			} else if k == 0 {
				dp[k][kk][0] = dp[k][kk-1][0] + matrix[k][kk]
				dp[k][kk][1] = matrix[k][kk]
			} else if kk == 0 {
				dp[k][kk][0] = matrix[k][kk]
				dp[k][kk][1] = dp[k-1][kk][1] + matrix[k][kk]
			} else {
				dp[k][kk][0] = dp[k][kk-1][0] + matrix[k][kk]
				dp[k][kk][1] = dp[k-1][kk][1] + matrix[k][kk]
			}
		}
	}
	return NumMatrix{dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	total := 0
	if row2-row1 <= col2-col1 {
		for ; row1 <= row2; row1++ {
			total += this.dp[row1][col2][0]
			if col1 > 0 {
				total -= this.dp[row1][col1-1][0]
			}
		}
	} else {
		for ; col1 <= col2; col1++ {
			total += this.dp[row2][col1][1]
			if row1 > 0 {
				total -= this.dp[row1-1][col1][1]
			}
		}
	}
	return total
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

/*
	用三维dp记录每个坐标的累积行列值
	这样只需要根据每个终点边上的行和列，都可以根据行首和行位计算出累计值
	在返回的时候，需要判断行的个数多还是列的个数多，然后取少的一边进行累加运算
*/
