package leetcode

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	flag := false
	var backtrace func(si, pi int, s, p string)
	backtrace = func(si, pi int, s, p string) {
		if si >= len(s) || pi >= len(p) || si < 0 || pi < 0 {
			if si == len(s) && pi == len(p) {
				flag = true
			}
			return
		}
		if !flag && si < len(s) && pi < len(p) {
			if pi+2 < len(p) && p[pi+2] == '*' {
				char := string(p[pi+1])
				count := 1
				for tsi := si + 1; tsi < len(s) && (s[tsi] == p[pi+1] || (p[pi+1] == '.')); tsi++ {
					count++
				}
				temp := ""
				for i := 0; i <= count; i++ {
					backtrace(si, pi, s, p[:pi+1]+temp+p[pi+3:])
					temp += char
				}
			} else if pi+1 < len(p) && p[pi+1] == '*' {
				char := string(p[pi])
				count := 1
				for tsi := si; tsi < len(s) && (s[tsi] == p[pi] || (p[pi] == '.')); tsi++ {
					count++
				}
				temp := ""
				for i := 0; i <= count; i++ {
					backtrace(si, pi, s, p[:pi]+temp+p[pi+2:])
					temp += char
				}
			} else if s[si] == p[pi] || p[pi] == '.' {
				backtrace(si+1, pi+1, s, p)
			}
		}
	}
	backtrace(0, 0, s, p)
	return flag
}

// @lc code=end

/*
	回溯 + 剪枝

	回溯：
		本质上就是当遍历到p字符串i索引时，当前字符为*,那么就替换 i-1 索引的字符 为 0-n 个
		为了少回头再遍历，直接在当前位置就 预先 判断后面第三个是不是为 * ，这样在进入下一轮判断的时候，就直接是 si+1 和 pi+1
		但是有特殊情况， 比如 s = "aa" , p = "a*"  这种情况就只能判断 后面第二个了。
		于是： 判断字符为  *  就有如下两种
			优先判断第三个 if  pi+2 < len(p) && p[pi+2] == '*'
			其次再判断第二个  if pi+1 < len(p) && p[pi+1] == '*'

		根据当前是判断的第三个还是第二个，直接就复制 当前的字符(第二个为*)，或者下一个字符(第三个为*)

	剪枝：
		就是减少 * 前面字符的复制次数，s和p的遍历索引相同
		两种条件下需要继续复制,count用于标识当前应该复制的次数
			判断p中出现 * 时的前一个字符的索引与对应的 s 的索引字符是否相同,且 s[索引] 不为 .
			s[对应的索引]  为 .

*/
