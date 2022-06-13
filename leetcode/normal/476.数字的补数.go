package leetcode

/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	binary := []int{}
	for num > 0 {
		binary = append(binary, num%2)
		num /= 2
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
