package leetcode

/*
 * @lc app=leetcode.cn id=1803 lang=golang
 *
 * [1803] 统计异或值在范围内的数对有多少
 */

// @lc code=start
func countPairs(nums []int, low int, high int) int {
	result := 0
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for high++; high > 0; high >>= 1 {
		nxt := map[int]int{}
		for x, c := range cnt {
			result += c * (high%2*cnt[(high-1)^x] - low%2*cnt[(low-1)^x])
			nxt[x>>1] += c
		}
		cnt = nxt
		low >>= 1
	}
	return result / 2
}

// @lc code=end

/*
参考灵神题解按照比特位1划分区间,减少异或XOR运算中的比特位
即可判断两个数字是否处于该区间中


暴力枚举运算超时
Time Limit Exceeded
50/63 cases passed (N/A)
直接枚举所有组合，然后依次 xor 运算
func countPairs(nums []int, low int, high int) int {
	val := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if val = nums[i] ^ nums[j]; val >= low && val <= high {
				result++
			}
		}
	}
	return result
}

*/
