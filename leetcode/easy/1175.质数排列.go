package easy

/*
 * @lc app=leetcode.cn id=1175 lang=golang
 *
 * [1175] 质数排列
 */

// @lc code=start
func numPrimeArrangements(n int) int {
	//厄拉多塞筛法
	ht := map[int]bool{}
	for i := 2; i <= n; i++ {
		ht[i] = true
	}
	temp := 0
	for i := 2; i <= n; i++ {
		temp = i * 2
		if ht[i] {
			for temp <= n {
				if ht[temp] {
					delete(ht, temp)
				}
				temp += i
			}
		}
	}
	primeCount := len(ht)
	NotPrimeCount := n - len(ht)
	temp1, temp2 := 1, 1
	var mod int = 1e9 + 7
	for primeCount > 0 {
		temp1 *= primeCount
		temp1 %= mod
		primeCount--
	}
	for NotPrimeCount > 0 {
		temp2 *= NotPrimeCount
		temp2 %= mod
		NotPrimeCount--
	}
	return temp1 * temp2 % mod
}

// @lc code=end
