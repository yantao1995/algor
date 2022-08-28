package leetcode

/*
 * @lc app=leetcode.cn id=793 lang=golang
 *
 * [793] 阶乘函数后 K 个零
 */

// @lc code=start
func preimageSizeFZF(k int) int {
	if k < 2 {
		return 5
	}
	//172 题 获得阶乘0的个数
	trailingZeroes := func(n int) int {
		val := 0
		for n >= 5 {
			n /= 5
			val += n
		}
		return val
	}
	var l, r int
	for l, r = 1, 1e10; l < r; {
		mid := (l + r + 1) >> 1
		if result := trailingZeroes(mid); result == k {
			return 5
		} else if result < k {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return 0
}

// @lc code=end

/*
	以2和5为质因子,得到10，那么可以得到count(2) >= count(5)
	所以只需要找到质因子5的数量就是0的数量，因为数字增加，0的个数必然增加。
	所以以5为临界点，可以得到个数要么是5，要么是0。如 5! ~ 9!
	即：只要找到一个数满足，即可直接返回5

	因为 k 的范围： 0 <= k <= 109 , 所以可以在该范围进行二分搜索

*/
