package leetcode

/*
 * @lc app=leetcode.cn id=1224 lang=golang
 *
 * [1224] 最大相等频率
 */

// @lc code=start
func maxEqualFreq(nums []int) int {
	result := 0
	maxCount := 0
	cm := map[int]int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	equals := 0
	maxCountSub1 := 0
	for k := range nums {
		cm[nums[k]]++
		if cm[nums[k]] > maxCount {
			maxCount = cm[nums[k]]
			maxCountSub1 = equals - 1
			equals = 1
		} else if cm[nums[k]] == maxCount {
			equals++
			maxCountSub1--
		} else if cm[nums[k]] == maxCount-1 {
			maxCountSub1++
		}

		if equals == len(cm) { //全相等
			if maxCount == 1 || equals == 1 { //全为1 或者只有1个数
				result = max(result, k+1)
			} else if k+1 < len(nums) { //可以随便加一个进来
				result = max(result, k+2)
			} else if equals == 2 {
				result = max(result, k)
			}
		} else if maxCountSub1 == len(cm)-1 { //最大的减一个即可以全相等
			result = max(result, k+1)
		} else if equals == len(cm)-1 && maxCount*equals == k { //1个为1，其他的全相等
			result = max(result, k+1)
		}
	}
	return result
}

// @lc code=end

/*
	面向测试用例编程
	前缀：从0-n遍历，走到k，那么k前面的均为前缀

	maxCount 单个数的最大重复次数
	cm 存储每个数的重复次数
	equals 与maxCount个数相同的数的个数
	maxCountSub1 与 maxCount-1个数相同的数的个数

	判断当前能否为结果的逻辑：
		1.如果所有数的个数全部相等：(因为k是从0开始，所以k+1表示当前前缀的个数)
			a.个数全为1，或者所有数相同，那么当前的所有数可以达成
			b.如果k+1还是数组元素，那么就可以随便加一个进来，然后删掉，就可以完成全部相等
			c.如果当前只有两个不相同的数，那么当前数的前一个数可以达成，其实和b重复了
		2.如果 maxCountSub1 的个数是不同数的个数的总数-1，那么maxCount的数-1，即可全部相等
		3.如果有一个数为1，其他的数全部相等，那么可以达成


	/
		下面的代码超时,但是需要统计 equals 和 maxCountSub1 的个数
		if cm[nums[k]] > maxCount {
			maxCount = cm[nums[k]]
		}
		equals = 0
		maxCountSub1 = 0
		for _, c := range cm {
			if c == maxCount {
				equals++
			} else if c == maxCount-1 {
				maxCountSub1++
			}
		}

		换用新逻辑： 减少 n*m 次遍历
		if cm[nums[k]] > maxCount {
			maxCount = cm[nums[k]]
			maxCountSub1 = equals - 1
			equals = 1
		} else if cm[nums[k]] == maxCount {
			equals++
			maxCountSub1--
		} else if cm[nums[k]] == maxCount-1 {
			maxCountSub1++
		}
	/

*/
