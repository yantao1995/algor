package leetcode

/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	result := make([]int, 0, len(mat)*len(mat[0]))
	leftUp := true
	for i, j, up := 0, 0, true; i < len(mat) && j < len(mat[0]); {
		result = append(result, mat[i][j])
		if i == len(mat)-1 && j == len(mat[i])-1 {
			break
		}
		if up {
			if i == 0 || j == len(mat[i])-1 {
				up = !up
				if i == 0 && j == len(mat[i])-1 {
					i++
					continue
				}
			}
			if leftUp {
				if i > 0 {
					i--
				}
				if j < len(mat)-1 {
					j++
				}
			} else { //TODO  左上部分 和右下部分分开判断
				if i > 0 {
					i--
				}
			}

		} else {
			if i == len(mat)-1 || j == 0 {
				up = !up
				if i == len(mat)-1 && j == 0 {
					j++
					continue
				}
			}
			if i < len(mat)-1 {
				i++
			}
			if j > 0 {
				j--
			}
		}
	}
	return result
}

// @lc code=end
