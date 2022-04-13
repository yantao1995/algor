package leetcode

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	isValid := false
	flagMap := map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9}
	var backtrace func(i int)
	backtrace = func(i int) {
		if i >= 9 || isValid {
			return
		}
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				flag := make([]bool, 10)
				for temp := 0; temp < 9; temp++ {
					if board[temp][j] != '.' {
						flag[flagMap[board[temp][j]]] = true
					}
					if board[i][temp] != '.' {
						flag[flagMap[board[i][temp]]] = true
					}
				}
				for round, ti, tj := 0, i/3, j/3; round < 3; round, ti, tj = round+1, ti+1, tj+1 {
					if board[ti][tj] != '.' {
						flag[flagMap[board[ti][tj]]] = true
					}
				}
				for m := byte('1'); m <= '9'; m++ {
					if !flag[flagMap[m]] {
						board[i][j] = m
						backtrace(i + 1)
						board[i][j] = '.'
					}
				}
			}
		}
	}
	backtrace(0)
}

// @lc code=end
