package leetcode

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	rtn, temp1, temp2, temp3 := 99999999999, 99999999999, 0, 0
	for i := 0; i+2 < len(nums); i++ {
		for j := i + 1; j+1 < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				temp2 = nums[i] + nums[j] + nums[k]
				temp3 = getAbsoultDiff(temp2, target)
				if temp3 < temp1 {
					rtn = temp2
					temp1 = temp3
				}
			}
		}
	}
	return rtn
}

func getAbsoultDiff(value1, value2 int) int {
	if value1 < value2 {
		return value2 - value1
	}
	return value1 - value2
}

// @lc code=end
