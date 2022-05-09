package leetcode

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if x == 1 || n == 0 {
		return 1
	}
	flag := n > 0
	if !flag {
		n = 0 - n
	}
	memo := map[int]float64{0: 1, 1: x}
	var binary func(n int) float64
	binary = func(n int) float64 {
		if _, ok := memo[n]; ok {
			return memo[n]
		}
		memo[n] = binary(n/2) * binary(n/2)
		if n&1 == 1 {
			memo[n] = memo[n] * x
		}
		return memo[n]
	}
	if flag {
		return binary(n)
	}
	return 1 / binary(n)
}

// @lc code=end

/*
	常规直接累乘，会超时，所以可以
	利用 2^10 = 2^5 * 2^5 来进行二分累乘，降低时间复杂度
	flag 记录当前是正值还是负值，用于结果取倒数
	memo记录已经计算过的数，避免重复计算
	然后递归进行累乘计算
*/
