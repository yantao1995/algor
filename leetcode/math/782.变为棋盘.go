package leetcode

/*
 * @lc app=leetcode.cn id=782 lang=golang
 *
 * [782] 变为棋盘
 */

// @lc code=start
func movesToChessboard(board [][]int) int {
	n, rowSum, colSum, rowDiff, colDiff := len(board), 0, 0, 0, 0
	//每个矩形框的四个顶点应该满足 4个0，4个1，2个0和2个1 即异或不为0
	for k := range board {
		for kk := range board[k] {
			if board[0][0]^board[k][kk]^
				board[0][kk]^board[k][0] == 1 {
				return -1
			}
		}
	}
	//计算每个行列的1和0的数量以及变换次数
	for i := 0; i < n; i++ {
		rowSum += board[0][i]
		colSum += board[i][0]
		if !(board[i][0] == i%2) {
			rowDiff++
		}
		if !(board[0][i] == i%2) {
			colDiff++
		}
	}

	//每行每列 1或0的数量应该是 n/2 或者 (n+1)/2
	if (rowSum < n/2 || rowSum > (n+1)/2) ||
		(colSum < n/2 || colSum > (n+1)/2) {
		return -1
	}

	//n为奇数
	if n%2 == 1 {
		if rowDiff%2 == 1 {
			rowDiff = n - rowDiff
		}
		if colDiff%2 == 1 {
			colDiff = n - colDiff
		}
	} else {
		if n-rowDiff < rowDiff {
			rowDiff = n - rowDiff
		}
		if n-colDiff < colDiff {
			colDiff = n - colDiff
		}
	}
	return (rowDiff + colDiff) / 2
}

// @lc code=end

/*
	参考  https://www.cnblogs.com/grandyang/p/9053705.html
*/
