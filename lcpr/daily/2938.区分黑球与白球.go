package lcpr

/*
 * @lc app=leetcode.cn id=2938 lang=golang
 * @lcpr version=30203
 *
 * [2938] 区分黑球与白球
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumSteps(s string) int64 {
	result := 0
	//数0的个数
	count0 := 0
	for i := range s {
		if s[i] == '0' {
			count0++
		}
	}
	//记录要移动的1的位置
	mIndex := map[int]bool{}
	for i := 0; i < count0; i++ {
		if s[i] == '1' {
			mIndex[i] = true
		}
	}
	//记录
	for i := count0; i < len(s); i++ {
		if s[i] == '0' {
			//找空余位置
			for k := range mIndex {
				result += i - k
				delete(mIndex, k)
				break
			}
		}
	}
	return int64(result)
}

// @lc code=end

/*
// @lcpr case=start
// "101"\n
// @lcpr case=end

// @lcpr case=start
// "100"\n
// @lcpr case=end

// @lcpr case=start
// "0111"\n
// @lcpr case=end

*/

/*
 先计算0的总个数，如果左边全部填充0，那么1就要挪动，记录要挪动的1的位置
 然后从大于等于count0的位置开始向后找0，填到对应需要挪动的1的位置，计算距离即可
*/
