package leetcode

/*
 * @lc app=leetcode.cn id=2379 lang=golang
 *
 * [2379] 得到 K 个黑块的最少涂色次数
 */

// @lc code=start
func minimumRecolors(blocks string, k int) int {
	min := len(blocks)
	for i := 0; i+k-1 < len(blocks); i++ {
		count := 0
		for j := 0; j < k; j++ {
			if blocks[i+j] == 'W' {
				count++
			}
		}
		if count < min {
			min = count
		}
	}
	return min
}

// @lc code=end

/*
	指定窗口大小为k，然后数W的个数即可
*/
