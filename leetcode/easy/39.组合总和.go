package easy

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	rtn := [][]int{}
	QuickSort39(candidates, 0, len(candidates)-1)
	for k := len(candidates) - 1; k >= 0; k-- {
		if candidates[k] == target {
			rtn = append(rtn, []int{candidates[k]})
		} else if candidates[k] > target {
			continue
		} else {
			tempSlice := []int{}
			tempIndexSlice := []int{}
			datas := [][]int{}
			getNums(candidates, tempSlice, tempIndexSlice, &datas, k, 0, target)
			if len(datas) > 0 {
				rtn = append(rtn, datas...)
			}
		}
	}
	return rtn
}

func getNums(candidates, tempSlice, tempIndexSlice []int, datas *[][]int, index, total, target int) {
	if index >= 0 { //还有结点 持续回溯
		if candidates[index]+total == target {
			total += candidates[index]
			tempSlice = append(tempSlice, candidates[index])
			tempIndexSlice = append(tempIndexSlice, index)
			ttt := make([]int, len(tempSlice)) // copy 防止下轮遍历改变append里面的内容
			copy(ttt, tempSlice)
			*datas = append(*datas, ttt)
			//可重复，再向下剪枝
			if len(tempSlice) > 1 && len(tempIndexSlice) > 1 { //
				total -= tempSlice[len(tempSlice)-1]
				tempSlice = tempSlice[:len(tempSlice)-1]
				tempIndexSlice = tempIndexSlice[:len(tempIndexSlice)-1]
				getNums(candidates, tempSlice, tempIndexSlice, datas, index-1, total, target)
			} else {
				return
			}
		} else if candidates[index]+total < target {
			total += candidates[index]
			tempSlice = append(tempSlice, candidates[index])
			tempIndexSlice = append(tempIndexSlice, index)
			getNums(candidates, tempSlice, tempIndexSlice, datas, index, total, target)
		} else {
			getNums(candidates, tempSlice, tempIndexSlice, datas, index-1, total, target)
		}
	} else { //剪枝
		if len(tempSlice) > 1 && len(tempIndexSlice) > 1 {
			cutIndex := 0
			total -= tempSlice[len(tempSlice)-1]
			tempSlice = tempSlice[:len(tempSlice)-1]
			cutIndex = tempIndexSlice[len(tempIndexSlice)-1]
			tempIndexSlice = tempIndexSlice[:len(tempIndexSlice)-1]
			getNums(candidates, tempSlice, tempIndexSlice, datas, cutIndex-1, total, target)
		} else {
			return
		}
	}
}

func QuickSort39(nums []int, low, high int) {
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
				nums[index], nums[j] = nums[j], nums[index]
				index = j
				sufftoprev = false
			}
			j--
		} else { //坑在后面，哨兵从前往后
			if nums[index] < nums[i] {
				nums[index], nums[i] = nums[i], nums[index]
				index = i
				sufftoprev = true
			}
			i++
		}
	}
	QuickSort39(nums, low, index-1)
	QuickSort39(nums, index+1, high)
}

// @lc code=end
