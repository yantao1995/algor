package leetcode

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	result := []string{}
	type local struct {
		x, y int
	}
	bitByte := [10]map[byte]bool{}
	for i := 0; i < 10; i++ {
		bitByte[i] = map[byte]bool{}
		for k := range words {
			if len(words[k]) > i {
				bitByte[i][words[k][i]] = true
			}
		}
	}
	router := map[local]bool{}
	distinct := map[string]bool{}
	var backtrace func(length int, lc local, log []byte)
	backtrace = func(length int, lc local, log []byte) {
		if lc.x < 0 || lc.x >= len(board) || lc.y < 0 ||
			lc.y >= len(board[lc.x]) || router[lc] {
			return
		}
		if length < 10 && bitByte[length][board[lc.x][lc.y]] {
			log = append(log, board[lc.x][lc.y])
			distinct[string(log)] = true
			router[lc] = true
			backtrace(length+1, local{lc.x + 1, lc.y}, log)
			backtrace(length+1, local{lc.x - 1, lc.y}, log)
			backtrace(length+1, local{lc.x, lc.y + 1}, log)
			backtrace(length+1, local{lc.x, lc.y - 1}, log)
			delete(router, lc)
		}
	}
	for k := range board {
		for kk := range board[k] {
			backtrace(0, local{k, kk}, []byte{})
		}
	}
	for k := range words {
		if distinct[words[k]] {
			result = append(result, words[k])
		}
	}
	return result
}

// @lc code=end

/*
	回溯
	local为每一个board的坐标点
	思路：从每一个起点,对每一个可用的单词向4个方向进行回溯，router去重来的路径
		 然后使用distinct记录每一个可达的单词，就能得到可以生成的所有单词
	剪枝：
		1.因为当前长度最大最大10，所以10以上长度的单词直接剪枝
		2.根据给定的结果集，记录每一位能出现的字符,然后回溯时，当前字符如果不存在结果集里面,也直接剪枝

*/
