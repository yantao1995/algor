package easy

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rowMap := [9]map[byte]bool{} //行
	colMap := [9]map[byte]bool{} //列
	suqMap := [9]map[byte]bool{} //方格
	// for k := 0; k < 9; k++ {

	// }
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[byte]bool)
		colMap[i] = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if j%3 == 0 && i%3 == 0 {
				suqMap[i/3*3+j/3] = make(map[byte]bool)
			}
			if board[i][j] != '.' {
				if !rowMap[i][board[i][j]] { // 行
					rowMap[i][board[i][j]] = true
				} else {
					return false
				}
				if !suqMap[i/3*3+j/3][board[i][j]] { //方格
					suqMap[i/3*3+j/3][board[i][j]] = true
				} else {
					return false
				}
			}
			if board[j][i] != '.' {
				if !colMap[i][board[j][i]] { //列
					colMap[i][board[j][i]] = true
				} else {
					return false
				}
			}
		}
	}
	return true
}

// @lc code=end
