package leetcode

/*
 * @lc app=leetcode.cn id=1640 lang=golang
 *
 * [1640] 能否连接形成数组
 */

// @lc code=start
func canFormArray(arr []int, pieces [][]int) bool {
	can := func(i, j int) bool {
		for m := 0; m < len(pieces[j]); m++ {
			if arr[i+m] != pieces[j][m] {
				return false
			}
		}
		return true
	}
	distinct := map[int]bool{}
lab:
	for i := 0; i < len(arr); {
		for j := 0; j < len(pieces); j++ {
			if !distinct[j] && can(i, j) {
				i += len(pieces[j])
				distinct[j] = true
				continue lab
			}
		}
		return false
	}
	return true
}

// @lc code=end

/*
	can 判断arr在i索引是否能匹配到 pieces[j]
	distinct 记录已经匹配到的 pieces 索引，避免重复匹配
	如果在当前 i 索引 未找到可以匹配的，那就不能匹配，否则继续匹配
*/
