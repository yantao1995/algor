package leetcode

/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */

// @lc code=start
func canTransform(start string, end string) bool {
	i, j := 0, 0
	for ; i < len(start) || j < len(end); i, j = i+1, j+1 {
		for i < len(start) && start[i] == 'X' {
			i++
		}
		for j < len(end) && end[j] == 'X' {
			j++
		}
		if i == len(start) || j == len(end) {
			break
		}
		if start[i] != end[j] ||
			(start[i] == 'L' && i < j) ||
			(start[i] == 'R' && i > j) {
			return false
		}
	}
	return i == j
}

// @lc code=end

/*
	转化性质可以理解为 L可以借着相邻的X 向左移动， R借着X向右移动
	判断第对应顺位的L和R所在位置是否无法移动
	如果start的第i个L左边没有X,而end的第j+1个L左边也没有X，并且位置与start的第i个相等，
		相当于两个L位置错开，并且无法移动，那么不可完成
	R同理，那么无法达成
*/
