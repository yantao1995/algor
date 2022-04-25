package leetcode

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	memo := make([][]bool, m)
	for k := range memo {
		memo[k] = make([]bool, n)
	}
	direct := 1
	for i, j := 0, 0; len(result) < m*n; {
		result = append(result, matrix[i][j])
		memo[i][j] = true
		switch direct {
		case 1:
			if j+1 < n && !memo[i][j+1] {
				j++
				continue
			}
			i++
			direct = 2
		case 2:
			if i+1 < m && !memo[i+1][j] {
				i++
				continue
			}
			j--
			direct = 3
		case 3:
			if j-1 >= 0 && !memo[i][j-1] {
				j--
				continue
			}
			i--
			direct = 4
		case 4:
			if i-1 >= 0 && !memo[i-1][j] {
				i--
				continue
			}
			j++
			direct = 1
		}

	}
	return result
}

// @lc code=end

/*
	memo 记录走过的位置
	direct记录当前方向
	按 向右， 向下， 向左， 向上的方向走
*/
