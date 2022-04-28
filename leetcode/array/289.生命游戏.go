package leetcode

/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

// @lc code=start
func gameOfLife(board [][]int) {
	var recursion func(i, j int)
	recursion = func(i, j int) {
		if j >= len(board[0]) {
			i++
			j = 0
			if i >= len(board) {
				return
			}
		}
		count := 0
		if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] == 1 {
			count++
		}
		if i-1 >= 0 && board[i-1][j] == 1 {
			count++
		}
		if i-1 >= 0 && j+1 < len(board[i]) && board[i-1][j+1] == 1 {
			count++
		}
		if j-1 >= 0 && board[i][j-1] == 1 {
			count++
		}
		if j+1 < len(board[i]) && board[i][j+1] == 1 {
			count++
		}
		if i+1 < len(board) && j-1 >= 0 && board[i+1][j-1] == 1 {
			count++
		}
		if i+1 < len(board) && board[i+1][j] == 1 {
			count++
		}
		if i+1 < len(board) && j+1 < len(board[i]) && board[i+1][j+1] == 1 {
			count++
		}
		recursion(i, j+1)
		if board[i][j] == 1 {
			if count < 2 || count > 3 {
				board[i][j] = 0
			}
		} else {
			if count == 3 {
				board[i][j] = 1
			}
		}

	}
	recursion(0, 0)
}

// @lc code=end

/*
	因为需要使用原地算法，所以在递归前记录下当前board[i][j]的值
	先判断，进入递归，最后再更新值
*/
