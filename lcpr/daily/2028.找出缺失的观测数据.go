package lcpr

/*
 * @lc app=leetcode.cn id=2028 lang=golang
 * @lcpr version=30202
 *
 * [2028] 找出缺失的观测数据
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func missingRolls(rolls []int, mean int, n int) []int {
	remainder := (len(rolls) + n) * mean
	for _, v := range rolls {
		remainder -= v
	}
	avg := float64(remainder) / float64(n)
	if avg > 6 || avg < 1 {
		return nil
	}
	result := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		result[i] = int(avg)
		total += result[i]
	}
	for i := 0; total != remainder; i++ {
		diff := 6 - result[i]
		if total+diff > remainder {
			diff = remainder - total
		}
		total += diff
		result[i] += diff
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,4,3]\n4\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,5,6]\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n6\n4\n
// @lcpr case=end

// @lcpr case=start
// [1]\n3\n1\n
// @lcpr case=end

*/

/*
	直接算平均数，然后向下取整，给每个数加进去。最后把多余的整数依次给数组按规则加进去
*/
