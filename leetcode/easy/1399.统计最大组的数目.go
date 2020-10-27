package easy

/*
 * @lc app=leetcode.cn id=1399 lang=golang
 *
 * [1399] 统计最大组的数目
 */

// @lc code=start
func countLargestGroup(n int) int {
	ht := map[int]int{}
	maxCount, count := 0, 0
	for n > 0 {
		temp := n
		sum := 0
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		ht[sum]++
		if ht[sum] > maxCount {
			maxCount = ht[sum]
		}
		n--
	}
	for k := range ht {
		if ht[k] == maxCount {
			count++
		}
	}
	return count
}

// @lc code=end
