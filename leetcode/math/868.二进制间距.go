package leetcode

/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(n int) int {
	count := 0
	for last, i := -1, 0; n > 0; n, i = n>>1, i+1 {
		if n&1 == 1 {
			if last > -1 && i-last > count {
				count = i - last
			}
			last = i
		}
	}
	return count
}

// @lc code=end

/*
	右移一位 与 1与运算，得到末尾是否为1，然后last记录上次为1的位置，与当前位置i做距离判断得到count
*/
