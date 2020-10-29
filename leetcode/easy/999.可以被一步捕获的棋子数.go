package leetcode

/*
 * @lc app=leetcode.cn id=999 lang=golang
 *
 * [999] 可以被一步捕获的棋子数
 */

// @lc code=start
func numRookCaptures(board [][]byte) int {
	result := 0
	for k := range board {
		for m := range board[k] {
			if board[k][m] == 'R' {
				for i := k + 1; i < len(board); i++ {
					if board[i][m] == 'R' || board[i][m] == 'B' {
						break
					}
					if board[i][m] == 'p' {
						result++
						break
					}
				}
				for i := k - 1; i >= 0; i-- {
					if board[i][m] == 'R' || board[i][m] == 'B' {
						break
					}
					if board[i][m] == 'p' {
						result++
						break
					}
				}
				for i := m - 1; i >= 0; i-- {
					if board[k][i] == 'R' || board[k][i] == 'B' {
						break
					}
					if board[k][i] == 'p' {
						result++
						break
					}
				}
				for i := m + 1; i < len(board[k]); i++ {
					if board[k][i] == 'R' || board[k][i] == 'B' {
						break
					}
					if board[k][i] == 'p' {
						result++
						break
					}
				}
			}
		}
	}
	return result
}

// @lc code=end
