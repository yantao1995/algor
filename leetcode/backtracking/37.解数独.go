package leetcode

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	isValid := false
	type ft struct {
		sgl, idx int //标识  -1 行  -2 列   , 00 第一个粗实线分隔  01 第二个...
	}
	distMap := map[ft]map[byte]bool{}
	for i := 0; i < 9; i++ { //初始化行列map
		distMap[ft{-1, i}] = map[byte]bool{}        //行
		distMap[ft{-2, i}] = map[byte]bool{}        //列
		distMap[ft{i / 3, i % 3}] = map[byte]bool{} //九宫格
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				distMap[ft{-1, i}][board[i][j]] = true
				distMap[ft{-2, j}][board[i][j]] = true
				distMap[ft{i / 3, j / 3}][board[i][j]] = true
			}
		}
	}

	var backtrace func(i, j int)
	backtrace = func(i, j int) {
		for i < 9 && j < 9 {
			if board[i][j] == '.' {
				break
			}
			j++
			if j == 9 {
				j = 0
				i++
			}

		}
		if isValid || i > 8 {
			isValid = true
			return
		}
		for m := byte('1'); m <= '9'; m++ {
			if !distMap[ft{-1, i}][m] && !distMap[ft{-2, j}][m] && !distMap[ft{i / 3, j / 3}][m] {
				board[i][j] = m
				distMap[ft{-1, i}][m] = true
				distMap[ft{-2, j}][m] = true
				distMap[ft{i / 3, j / 3}][m] = true
				backtrace(i, j)
				if isValid {
					return
				}
				board[i][j] = '.'
				distMap[ft{-1, i}][m] = false
				distMap[ft{-2, j}][m] = false
				distMap[ft{i / 3, j / 3}][m] = false
			}
		}
	}
	backtrace(0, 0)
}

// @lc code=end

/*
	type ft struct {
			sgl, idx int //标识  -1 行  -2 列   , 00 第一个粗实线分隔  01 第二个...
	}

	distMap 分别记录 0-8 (行，列，九宫格)  存在的 '1'-'9' 字符
	sgl 为 -1 标识当前为行    为 -2 标识当前为列， 九宫格按  从左到右，从上到下，分别标识为 0,0   0,1  0,2  1,0  1,1 1,2  2,0  2,1 2,2
	然后每次回溯 找到当前字符为 '.' ，分别从 '1'-'9' 分别尝试，每个回溯的函数体 都只处理 一个 '.'

	isValid 用来标识某一个回溯已经遍历过了整个 方格，  i>8 标识已经走出全部方格

	在接下来 需要回溯为 '.' 的地方进行跳出，防止数被改回去

*/
