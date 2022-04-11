package leetcode

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	nums := make([][]int, n)
	for k := range nums {
		nums[k] = make([]int, n)
	}
	i, j := 0, -1
	for seq := 1; seq <= n*n; seq++ {
		if j+1 < n && nums[i][j+1] == 0 {
			nums[i][j+1] = seq
			if j+2 < n && nums[i][j+2] == 0 {
				j++
			}
		} else if i+1 < n && nums[i+1][j] == 0 {
			nums[i+1][j] = seq
			if i+2 < n && nums[i+2][j] == 0 {
				i++
			}
		} else if j-1 >= 0 && nums[i][j-1] == 0 {
			nums[i][j-1] = seq
			if j-2 < n && nums[i][j-2] == 0 {
				j--
			}
		} else if i-1 >= 0 && nums[i-1][j] == 0 {
			nums[i-1][j] = seq
			if i-2 < n && nums[i-2][j] == 0 {
				i--
			}
		}
	}
	return nums
}

// @lc code=end
