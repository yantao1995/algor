package leetcode

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	ds, de := len(matrix), len(matrix[0])
	getNextSide := func(i, j, ds, de int) (int, int) {
		if i+1 < ds {
			i++
		}
		if j+1 < de {
			j++
		}
		return i, j
	}
	var scan func(i, j, ds, de int) bool
	scan = func(i, j, ds, de int) bool {
		tempI, tempJ := i, j
		for i < ds || j < de {
			if target == matrix[i][j] {
				return true
			} else if target > matrix[i][j] {
				if ii, jj := getNextSide(i, j, ds, de); ii == i && jj == j {
					return false
				} else {
					i, j = ii, jj
				}
			} else {
				if tempI == i && tempJ == j {
					return false
				}
				if scan(i, tempJ, ds, j) || scan(tempI, j, i, de) { //对角线斜边搜索，分块搜索
					return true
				}
				return false
			}
		}
		return false
	}
	return scan(0, 0, ds, de)
}

// @lc code=end

/*
	分治法
	因为二维矩阵有序递增，最优应该是使用斜边来进行二分查找，但是由于比较复杂，此处采用顺序斜边递增,
	即 [0,0],[1,1],[2,2]....的顺序来查找
	遇到比当前target大的值的时候，此时 i,j 即为边界，然后将剩余部分切为两块矩形，然后分别进入两块矩形安装上述步骤进行查找
*/
