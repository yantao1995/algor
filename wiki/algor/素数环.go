package algor

import "fmt"

/*
针对一个整数n，把从1到n的数字无重复的排列成环，且使每相邻两个数(包括首尾)的和都为素数，则称为素数环。
*/

func 素数环(n int) []int {
	//找到所有质数
	ht := map[int]bool{}
	for i := 2; i < 2*n; i++ { //因为 n 和  n-1 的和最大是 2n-1
		ht[i] = true
	}
	temp := 0
	for i := 2; i < 2*n; i++ {
		temp = i * 2
		if ht[i] {
			for temp < 2*n {
				if ht[temp] {
					delete(ht, temp)
				}
				temp += i
			}
		}
	}

	fmt.Println("2n-1的素数集合:", ht)
	// //构造 环形数组
	ring := make([]int, n)
	//记录某个位置的值是否被使用
	used := make([]bool, n+1)
	find := false //是否找到

	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n-1 {
			if ht[ring[n-1]+ring[0]] {
				find = true
			}
			return
		}
		//从未使用的数中选择
		for j := 1; j <= n; j++ {
			if find {
				return
			}
			if ht[ring[idx]+j] &&
				!used[j] { //相加是素数 并且未使用
				used[j] = true

				ring[idx+1] = j
				backtrace(idx + 1)

				used[j] = false
			}
		}

	}
	ring[0] = 1
	used[1] = true
	backtrace(0)
	return ring
}
