package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=556 lang=golang
 *
 * [556] 下一个更大元素 III
 */

// @lc code=start
func nextGreaterElement(n int) int {
	bits := []int{}
	for n > 0 {
		bits = append(bits, n%10)
		n /= 10
	}
	for i := 1; i < len(bits); i++ {
		tempIndex := 0
		find := false
		for j := 0; j < i; j++ {
			if bits[j] > bits[i] {
				if !find || bits[tempIndex] > bits[j] {
					tempIndex = j
				}
				find = true
			}
		}
		if find {
			bits[tempIndex], bits[i] = bits[i], bits[tempIndex]
			sort.Slice(bits[:i], func(i, j int) bool {
				return bits[j] < bits[i]
			})
			result := 0
			for j, m := 0, 1; j < len(bits); j, m = j+1, m*10 {
				result += bits[j] * m
			}
			if result > 1<<31-1 {
				break
			}
			return result
		}
	}
	return -1
}

// @lc code=end

/*
	思路：要想是最小的大于n的数，那么交换权值越低的数，值越小，然后再将剩下的数升序排列

	先按位取出各个位置的数字
	那么 循环i 从数字n的末尾依次向前扫描，然后 循环j 向末尾扫描,
	如果j中能找到一个比i大的数，那么就在该轮i中处理，因为要找最小的，所以还得找一个大于当前i，但是最小的数
	然后将 i  j 替换之后，对 i 之后的数做升序处理，即得到当前数字


*/
