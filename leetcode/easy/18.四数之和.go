package leetcode

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	QuickSort18(nums, 0, len(nums)-1)
	rtn := [][]int{}
	apdFlag := true
	if (target > 0 && nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3]+nums[len(nums)-4] < target) ||
		(target < 0 && nums[0]+nums[1]+nums[2]+nums[3] > target) {
		return nil
	}
	for i := 0; i+3 < len(nums); i++ {
		for j := i + 1; j+2 < len(nums); j++ {
			for k := j + 1; k+1 < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						apdFlag = true
						for kk := range rtn {
							if nums[i] == rtn[kk][0] && nums[j] == rtn[kk][1] &&
								nums[k] == rtn[kk][2] && nums[l] == rtn[kk][3] {
								apdFlag = false
								break
							}
						}
						if apdFlag {
							rtn = append(rtn, []int{nums[i], nums[j], nums[k], nums[l]})
						}
						break
					}
				}
			}
		}
	}
	return rtn
}

func QuickSort18(nums []int, low, high int) {
	if low >= high {
		return
	}
	i := low + 1
	j := high
	sufftoprev := true
	index := low
	for i <= j {
		if sufftoprev { //坑在前面，哨兵从后往前
			if nums[index] > nums[j] {
				exchangeIJ18(nums, index, j)
				index = j
				sufftoprev = false
			}
			j--
		} else { //坑在后面，哨兵从前往后
			if nums[index] < nums[i] {
				exchangeIJ18(nums, index, i)
				index = i
				sufftoprev = true
			}
			i++
		}
	}
	QuickSort18(nums, low, index-1)
	QuickSort18(nums, index+1, high)
}
func exchangeIJ18(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func getAbsoult(value int) int {
	if value < 0 {
		return 0 - value
	}
	return value
}

// @lc code=end
