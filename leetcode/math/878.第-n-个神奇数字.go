package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=878 lang=golang
 *
 * [878] 第 N 个神奇数字
 */

// @lc code=start
func nthMagicalNumber(n int, a int, b int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	//最大公约数 辗转相除法
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := a / gcd(a, b) * b
	return sort.Search(min(a, b)*n, func(i int) bool {
		return i/a+i/b-i/lcm >= n
	}) % (1e9 + 7)
}

// @lc code=end

/*
参考官方题解 ： 使用search 二分


暴力超时 44/70
func nthMagicalNumber(n int, a int, b int) int {
	mod := int(1e9 + 7)
	for i := 1; ; i++ {
		if i%a == 0 || i%b == 0 {
			n--
			if n == 0 {
				return i % mod
			}
		}
	}
}
*/
