package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=899 lang=golang
 *
 * [899] 有序队列
 */

// @lc code=start
func orderlyQueue(s string, k int) string {
	bts := []byte(s)
	if k > 1 {
		sort.Slice(bts, func(i, j int) bool {
			return bts[i] < bts[j]
		})
		return string(bts)
	}
	ans := s
	for i := 0; i < len(bts); i++ {
		bts = append(bts[1:], bts[:1]...)
		if ans > string(bts) {
			ans = string(bts)
		}
	}
	return ans
}

// @lc code=end

/*
	k>1 时，总能根据替换字符，找出任意序列的字符串，所以只需要排序即可
	k==1时，只能当做一个循环队列，找到一个最小的起点
*/
