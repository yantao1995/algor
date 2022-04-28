package leetcode

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	count := 0
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			count++
		}
	}
	return count
}

// @lc code=end

/*
	位运算，按位 与 1得到当前位是否为1
	然后累加
*/
