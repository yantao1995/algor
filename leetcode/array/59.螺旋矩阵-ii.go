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
	direct := 1 //1 右，2下 3左 4上
	for seq := 1; seq <= n*n; seq++ {
		switch direct {
		case 1:
			if j+1 < n && nums[i][j+1] == 0 {
				nums[i][j+1] = seq
				j++
			} else {
				direct = 2
				seq--
			}
		case 2:
			if i+1 < n && nums[i+1][j] == 0 {
				nums[i+1][j] = seq
				i++
			} else {
				direct = 3
				seq--
			}
		case 3:
			if j-1 >= 0 && nums[i][j-1] == 0 {
				nums[i][j-1] = seq
				j--
			} else {
				direct = 4
				seq--
			}
		case 4:
			if i-1 >= 0 && nums[i-1][j] == 0 {
				nums[i-1][j] = seq
				i--
			} else {
				direct = 1
				seq--
			}
		}
	}
	return nums
}

// @lc code=end
