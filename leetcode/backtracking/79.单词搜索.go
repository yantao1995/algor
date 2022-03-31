package leetcode

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	dp := make([][]bool, len(board))
	for k := range dp {
		dp[k] = make([]bool, len(board[k]))
	}
	var backtrack func(i, j, index int) bool
	backtrack = func(i, j, index int) bool {
		if i >= 0 && i < len(board) &&
			j >= 0 && j < len(board[i]) &&
			!dp[i][j] && word[index] == board[i][j] {
			dp[i][j] = true
			if len(word)-1 == index {
				return true
			}
			if backtrack(i-1, j, index+1) ||
				backtrack(i+1, j, index+1) ||
				backtrack(i, j-1, index+1) ||
				backtrack(i, j+1, index+1) {
				return true
			}
			dp[i][j] = false
		}
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] &&
				backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end

/*
	回溯算法，用dp二维数组记录每个值是否遍历过
	因为可以通过相邻单元格，每次向4个方向遍历，遇到边界或者遍历过或者不是当前索引的字符，就退出。
*/
