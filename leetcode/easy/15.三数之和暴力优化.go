package easy

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	QuickSort(nums, 0, len(nums)-1)
	rtn := [][]int{}
	apdFlag := true
	for i := 0; nums[i] <= 0 && i < len(nums)-2; i++ {
		if 0-nums[i] > nums[len(nums)-1]+nums[len(nums)-2] {
			continue
		}
		for j := i + 1; j+1 < len(nums); j++ {
			if nums[j] < 0 {
				if nums[i]+nums[j] < 0-nums[len(nums)-1] {
					continue
				}
			} else {
				if 0-nums[i] > nums[j]+nums[len(nums)-1] {
					continue
				}
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					apdFlag = true
					for kk := range rtn {
						if nums[i] == rtn[kk][0] && nums[j] == rtn[kk][1] &&
							nums[k] == rtn[kk][2] {
							apdFlag = false
							break
						}
					}
					if apdFlag {
						rtn = append(rtn, []int{nums[i], nums[j], nums[k]})
					}
					break
				}
			}
		}
	}

	return rtn
}

func QuickSort(nums []int, low, high int) {
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
				exchangeIJ(nums, index, j)
				index = j
				sufftoprev = false
			}
			j--
		} else { //坑在后面，哨兵从前往后
			if nums[index] < nums[i] {
				exchangeIJ(nums, index, i)
				index = i
				sufftoprev = true
			}
			i++
		}
	}
	QuickSort(nums, low, index-1)
	QuickSort(nums, index+1, high)
}
func exchangeIJ(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// @lc code=end
