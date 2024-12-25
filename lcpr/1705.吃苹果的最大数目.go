package lcpr

/*
 * @lc app=leetcode.cn id=1705 lang=golang
 * @lcpr version=30204
 *
 * [1705] 吃苹果的最大数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func eatenApples(apples []int, days []int) int {
	result := 0
	mCount := map[int]int{} //map[i索引:个数]
	mDay := map[int]int{}   //map[i索引:天数]
	find := func() {
		findFlag := false
		for k := range mCount {
			if mCount[k] > 0 && mDay[k] > 0 {
				findFlag = true
			}
			mCount[k]--
			mDay[k]--
			if mCount[k] <= 0 || mDay[k] <= 0 {
				delete(mCount, k)
			}
		}
		if findFlag {
			result++
		}
	}
	for i := 0; i < len(apples); i++ {
		mCount[i] = apples[i]
		mDay[i] = days[i]
	}
	for i := 0; len(mCount) > 0; i++ {
		find()
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,5,2]\n[3,2,1,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,0,0,0,0,2]\n[3,0,0,0,0,2]\n
// @lcpr case=end

*/
