package lcpr

/*
 * @lc app=leetcode.cn id=2972 lang=golang
 * @lcpr version=30204
 *
 * [2972] 统计移除递增子数组的数目 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func incremovableSubarrayCount(nums []int) int64 {
	total := 1 //全部删除的个数
	pre := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] <= nums[i-1] {
			break
		}
		pre = i
		total++ //只剩下左边的个数
	}
	if pre+1 == len(nums) { //数组本身递增，直接公式计算
		return int64(len(nums) * (len(nums) + 1) / 2)
	}
	for suf := len(nums) - 1; suf >= 0; suf-- { //右指针依次取1~n个
		if suf < len(nums)-1 && nums[suf] >= nums[suf+1] {
			break
		}
		i := 0 //左指针i记录当前左边能到的最大索引位置
		for i < suf {
			if nums[i] >= nums[suf] || (i > 0 && nums[i] <= nums[i-1]) {
				i-- //不包含当前位置i
				break
			}
			i++
		}
		if i >= 0 { //至少左边第一个元素满足条件
			total += i + 1 //左边剩余[1:i]个 (i从0开始,所以i+1)
		}
		total++ // 只剩下右边
	}
	return int64(total)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [6,5,7,8]\n
// @lcpr case=end

// @lcpr case=start
// [8,7,6,6]\n
// @lcpr case=end

*/

/*
优化思路：
	双指针分别记录
	左边从左至右连续递增的数组[0:pre]，
	右边从右至左连续递减的数组[suf:len(nums)-1]。
	然后依次匹配个数即可。
	第一个for记录左边依次选择递增，然后右边全部删除的情况。
	第二个for记录右边依次选择递减，然后左边依次选择[0:i]个  + 左边全布删除 的情况





Time Limit Exceeded
827/839 cases passed (N/A)

代码同 2970.统计移除递增子数组的数目 I
func incremovableSubarrayCount(nums []int) int64 {
	total := int64(0)
	isIncrease := func(start, end int) bool {
		gap := 1
		last := -1
		for i := 0; i < len(nums) && gap < 2; i++ {
			if i == start {
				i = end
				continue
			}
			if nums[i] > last {
				last = nums[i]
			} else {
				gap++
			}
		}
		return gap < 2
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if isIncrease(i, j) {
				total++
			}
		}
	}
	return total
}
*/
