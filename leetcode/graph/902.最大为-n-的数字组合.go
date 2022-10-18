package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=902 lang=golang
 *
 * [902] 最大为 N 的数字组合
 */

// @lc code=start
func atMostNGivenDigitSet(digits []string, n int) int {
	nSli := strings.Split(strconv.Itoa(n), "")
	result := 0
	last := len(digits)
	for l := 1; l < len(nSli); l++ {
		result += last
		last *= len(digits)
	}

	var dfs func(lastLess bool, idx int) int
	dfs = func(lastLess bool, idx int) int {
		if idx < len(nSli) {
			if lastLess {
				return len(digits) * dfs(lastLess, idx+1)
			}
			result := 0
			for j := 0; j < len(digits); j++ {
				if digits[j] > nSli[idx] {
					break
				}
				result += dfs(digits[j] < nSli[idx], idx+1)
			}
			return result
		}
		return 1
	}
	for i := 0; i < len(digits); i++ {
		if digits[i] > nSli[0] {
			break
		}
		result += dfs(digits[i] < nSli[0], 1)
	}
	return result
}

// @lc code=end

/*
	dfs
	先拿到n的位数，若此时n的位数为m，当使用 digits 的任意位组合,得到长度小于 m 位的数时，都可以满足条件
	那么此时分两种情况讨论：
		1. 位数	< m  , digits 全排列枚举所有可能
		2. 位数 = m  , 依次从 n 的权重最高的位向后计算。
			a.若当前位能从 digits 找到小于的数，那么小于的数往后，取任意数都可以满足，即后面直接全排列组合。
			b.若当前位和从 digits 找到的数相等，那么此时，紧接着后面的第一个位置只能找小雨等于当前位的数
			  若找到小于的数，继续执行a
			  若找到等于的数，继续执行b

	由条件2可以看到，此处使用递归最合适
*/
