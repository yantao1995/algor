package leetcode

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
// case17时超时
// func countPrimes(n int) int {
// 	count := 0
// 	n--
// 	for n > 1 {
// 		for i := 2; i < n; i++ {
// 			if n%i == 0 {
// 				goto A
// 			}
// 		}
// 		count++
// 	A:
// 		n--
// 	}
// 	return count
// }
func countPrimes(n int) int {
	//厄拉多塞筛法
	ht := map[int]bool{}
	for i := 2; i < n; i++ {
		ht[i] = true
	}
	temp := 0
	for i := 2; i < n; i++ {
		temp = i * 2
		if ht[i] {
			for temp < n {
				if ht[temp] {
					delete(ht, temp)
				}
				temp += i
			}
		}
	}
	return len(ht)
}

// @lc code=end
