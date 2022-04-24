package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	direct := 1
	for i, j, mm, nn := 0, 0, m-1, n-1; len(result) < m*n; {
		fmt.Println(" ___", direct, " ___")
		switch direct {
		case 1:
			for ti := nn; ti > 0; ti-- {
				fmt.Println(i, j)
				result = append(result, matrix[i][j])
				j++
			}
			direct = 2
		case 2:
			for ti := mm; ti > 0; ti-- {
				fmt.Println(i, j)
				result = append(result, matrix[i][j])
				i++
			}
			direct = 3
			mm--
		case 3:
			for ti := nn; ti > 0; ti-- {
				fmt.Println(i, j)
				result = append(result, matrix[i][j])
				j--
			}
			direct = 4
			nn--
		case 4:
			for ti := mm; ti > 0; ti-- {
				fmt.Println(i, j)
				result = append(result, matrix[i][j])
				i--
			}
			direct = 1
		}
	}
	return result
}

// @lc code=end
