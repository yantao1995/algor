package leetcode

/*
 * @lc app=leetcode.cn id=1653 lang=golang
 *
 * [1653] 使字符串平衡的最少删除次数
 */

// @lc code=start
func minimumDeletions(s string) int {
	prev, tail := make([]int, len(s)), make([]int, len(s))
	if s[0] == 'b' {
		prev[0] = 1
	}
	if s[len(s)-1] == 'a' {
		tail[len(s)-1] = 1
	}
	for i := 1; i < len(s); i++ {
		prev[i] = prev[i-1]
		if s[i] == 'b' {
			prev[i]++
		}
		tail[len(s)-1-i] = tail[len(s)-i]
		if s[len(s)-1-i] == 'a' {
			tail[len(s)-1-i]++
		}
	}
	min := len(s)
	for i := 0; i < len(s); i++ {
		if temp := prev[i] + tail[i] - 1; min > temp {
			min = temp
		}
	}
	return min
}

// @lc code=end

/*
	字符串平衡，即前面有且仅有a,后面有且仅有b
	那么可以分别从后往前和从前往后记录当前b的个数，a的个数
	因为要达到平衡状态，前面的b必须变成a，后面a必须变成b
	"aababbab"
	得到
	prev:[0 0 1 1 2 3 3 4]
	tail:[4 3 2 2 1 1 1 0]
	然后当前字符如果转换成另一个字符，得到转换次数为: prev[i]+tail[i]-1
	然后min记录最小次数即可
*/
