package leetcode

/*
 * @lc app=leetcode.cn id=764 lang=golang
 *
 * [764] 最大加号标志
 */

// @lc code=start
func orderOfLargestPlusSign(n int, mines [][]int) int {
	m := map[[2]int]bool{}
	temp := [2]int{}
	for _, mine := range mines {
		temp[0], temp[1] = mine[0], mine[1]
		m[temp] = true
	}
	dpUp, dpDown, dpLeft, dpRight := make([][]int, n), make([][]int, n), make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		dpUp[i] = make([]int, n)
		dpDown[i] = make([]int, n)
		dpLeft[i] = make([]int, n)
		dpRight[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dpLeft[i][j] = 0
			dpRight[i][n-j-1] = 0
			dpUp[i][j] = 0
			dpDown[n-i-1][j] = 0

			temp[0], temp[1] = i, j
			if !m[temp] {
				dpUp[i][j] = 1
				dpLeft[i][j] = 1
				if i > 0 {
					dpUp[i][j] += dpUp[i-1][j]
				}
				if j > 0 {
					dpLeft[i][j] += dpLeft[i][j-1]
				}
			}

			temp[0], temp[1] = n-i-1, j
			if !m[temp] {
				dpDown[n-i-1][j] = 1
				if i > 0 {
					dpDown[n-i-1][j] += dpDown[n-i][j]
				}
			}

			temp[0], temp[1] = i, n-j-1
			if !m[temp] {
				dpRight[i][n-j-1] = 1
				if j > 0 {
					dpRight[i][n-j-1] += dpRight[i][n-j]
				}
			}
		}
	}

	mins := func(a int, b ...int) int {
		for k := range b {
			if b[k] < a {
				a = b[k]
			}
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result = max(result, mins(dpUp[i][j], dpDown[i][j], dpLeft[i][j], dpRight[i][j]))
		}
	}

	return result
}

// @lc code=end

/*
	动态规划 dp[i]：到当前点的最大连续长度
	把矩形看成是4个方向上的dp数组，记录一个点的 上下左右 4个方向到当前的最大长度
	然后4个方向取最小的一个值就是当前的阶数
*/
