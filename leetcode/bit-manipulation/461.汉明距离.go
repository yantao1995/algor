package leetcode

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	result := x ^ y
	count := 0
	for result > 0 {
		if result&1 == 1 {
			count++
		}
		result >>= 1
	}
	return count
}

// @lc code=end

/*
	位运算规则:

	p	q	p&q	p|q	p^q
	0	0	0	0	0
	0	1	0	1	1
	1	1	1	1	0
	1	0	0	1	1

	使用异或 ^ 运算,相同为0，不同为1，来得到结果值
	比如  1011 ^ 1111  = 100 ， 此时 1的位置就是不同的值
	然后 判断 result&1 == 1 表示判断末尾是否为1 ，效果和 result%2 == 1 相同
	然后 使用移位运算符来对值进行除以二，即 result >>= 1 表一右移一位，效果和  result /= 2 相同
*/
