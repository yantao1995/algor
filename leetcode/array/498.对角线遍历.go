package leetcode

/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	result := make([]int, 0, len(mat)*len(mat[0]))
	up := true //斜向上
	for i, j := 0, 0; len(result) < cap(result); {
		result = append(result, mat[i][j])
		if up {
			if i == 0 || j == len(mat[0])-1 {
				up = !up
				if i == 0 && j < len(mat[i])-1 {
					j++
				} else {
					i++
				}
			} else {
				i--
				j++
			}
		} else {
			if i == len(mat)-1 || j == 0 {
				up = !up
				if j == 0 && i < len(mat)-1 {
					i++
				} else {
					j++
				}
			} else {
				i++
				j--
			}
		}
	}
	return result
}

// @lc code=end

/*
  up 记录当前的运行状态，只有在遇到边框的时候，才会变换运行状态
  从左上到右下的层次遍历，如果遇到一条边到了边框，变换方向，并且判断边界，然后移动i或j
*/
