package leetcode

/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
//取模。 1-59  2-58  组合总数  0，30 单独处理
func numPairsDivisibleBy60(time []int) int {
	ht := map[int]int{}
	for i := 0; i < len(time); i++ {
		ht[time[i]%60]++
	}
	result := 0
	i, j := 1, 59
	for i < j {
		if ht[i] > 0 && ht[j] > 0 {
			result += ht[i] * ht[j]
		}
		i++
		j--
	}
	if ht[0] > 0 {
		for m := ht[0] - 1; m > 0; m-- {
			result += m
		}
	}
	if ht[30] > 0 {
		for m := ht[30] - 1; m > 0; m-- {
			result += m
		}
	}
	return result
}

// @lc code=end
