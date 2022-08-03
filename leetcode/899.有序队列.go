package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=899 lang=golang
 *
 * [899] 有序队列
 */

// @lc code=start
func orderlyQueue(s string, k int) string {
	bts := []byte(s)
	minTail := byte('z')
	findMin := func() {
		for kk := k; kk < len(bts); kk++ {
			if bts[kk] < minTail {
				minTail = bts[kk]
			}
		}
	}
	temp := byte('z')
lab:
	minTail = byte('z')
	findMin()
	sort.Slice(bts[:k], func(i, j int) bool {
		return bts[i] < bts[j]
	})
	for i := 0; i < k; i++ {
		if bts[i] > minTail {
			temp = bts[i]
			bts = append(bts[:i], bts[i+1:]...)
			bts = append(bts, temp)
			goto lab
		}
	}
	return string(bts)
}

// @lc code=end
