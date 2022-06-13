package leetcode

/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(L int, R int) int {
	//厄拉多塞筛法
	ht := map[int]bool{}
	for i := 2; i < 32; i++ { //2进制 保存32内的质数
		ht[i] = true
	}
	temp := 0
	for i := 2; i < 32; i++ {
		temp = i * 2
		if ht[i] {
			for temp < 32 {
				if ht[temp] {
					delete(ht, temp)
				}
				temp += i
			}
		}
	}
	// ht 保存了 2-
	result := 0
	for L <= R {
		count := getNumberBits762(L)
		if ht[count] {
			result++
		}
		L++
	}
	return result
}

func getNumberBits762(n int) int {
	count := 0
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}

// @lc code=end
