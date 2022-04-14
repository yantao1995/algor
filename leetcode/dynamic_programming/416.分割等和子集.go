package leetcode

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	total := 0
	for k := range nums {
		total += nums[k]
	}
	if total%2 == 1 || len(nums) < 2 {
		return false
	}
	half := total / 2
	can := make([]bool, half+1)
	for k := range nums {
		for j := half; j > 0; j-- {
			if can[j] && j+nums[k] <= half {
				can[j+nums[k]] = true
			}
		}
		if nums[k] <= half {
			can[nums[k]] = true
		}
	}
	return can[total/2]
}

// @lc code=end

/*
	两个相等的的子集，即当前数组和的一半
	数组和为奇数 或者数组长度小于2均不可能完成
	此处只需要获得数组能组成 <= half 的所有可能
	定义half+1 长度的标志数组，为了取下标 can[half]，所以长度为half+1

	状态转移：如果当前值可以得到，即 can[j] ，那么 j + nums[k] 也可以得到，即  can[j + nums[k]],
	但由于需要找到值 half ， 所以 只需要  <= half 的所有值即可。

	每一个 can 数组的下标表示能否由 nums 数组中的元素相加得到
	如果当  can[half] =true 时，表示能得到 half 的值
*/
