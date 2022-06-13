package leetcode

/*
 * @lc app=leetcode.cn id=1560 lang=golang
 *
 * [1560] 圆形赛道上经过次数最多的扇区
 */

// @lc code=start
func mostVisited(n int, rounds []int) []int {
	ht := map[int]int{}
	max := 0
	for i := 1; i < len(rounds); i++ {
		if rounds[i] > rounds[i-1] {
			for rounds[i-1] < rounds[i] {
				ht[rounds[i-1]]++
				rounds[i-1]++
			}
		} else {
			for rounds[i-1] <= n {
				ht[rounds[i-1]]++
				rounds[i-1]++
			}
			tempI := rounds[i] - 1
			for tempI > 0 {
				ht[tempI]++
				tempI--
			}
		}
	}
	ht[rounds[len(rounds)-1]]++
	for k := range ht {
		if ht[k] > max {
			max = ht[k]
		}
	}
	result := []int{}
	for k := 1; k <= n; k++ {
		if max == ht[k] {
			result = append(result, k)
		}
	}
	return result
}

// @lc code=end
