package leetcode

/*
 * @lc app=leetcode.cn id=1009 lang=golang
 *
 * [1009] 十进制整数的反码
 */

// @lc code=start
func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	binary := []int{}
	for N > 0 {
		binary = append(binary, N%2)
		N /= 2
	}
	for k := range binary {
		if binary[k] == 1 {
			binary[k] = 0
		} else {
			binary[k] = 1
		}
	}
	total := 0
	weight := 1
	for k := range binary {
		if binary[k] == 1 {
			total += weight
		}
		weight *= 2
	}
	return total
}

// @lc code=end
