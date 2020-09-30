package leetcode

/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	prev, tail := make([]int, len(nums)), make([]int, len(nums))
	prev[0], tail[0] = nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		prev[i] = prev[i-1] + nums[i]
		tail[i] = tail[i-1] + nums[len(nums)-i-1]
	}
	for k := range prev {
		if prev[k] == tail[len(nums)-1-k] {
			return k
		}
	}
	return -1
}

// func pivotIndex(nums []int) int {
// 	prevIndex, tailIndex, prevTotal, tailTotal := -1, len(nums), 0, 0
// 	for tailIndex-prevIndex > 2 {
// 		if prevTotal < tailTotal {
// 			prevIndex++
// 			prevTotal += nums[prevIndex]
// 		} else {
// 			tailIndex--
// 			tailTotal += nums[tailIndex]
// 		}
// 		fmt.Println("prevIndex, tailIndex, prevTotal, tailTotal", prevIndex, tailIndex, prevTotal, tailTotal)
// 		for tailIndex-prevIndex > 3 &&
// 			((prevIndex > -1 && nums[prevIndex] == 0) || (tailIndex < len(nums) && nums[tailIndex] == 0)) {
// 			fmt.Println("+++++++++++= prevIndex, tailIndex, prevTotal, tailTotal", prevIndex, tailIndex, prevTotal, tailTotal)
// 			if prevIndex > -1 && nums[prevIndex] == 0 {
// 				prevIndex++
// 			} else if tailIndex < len(nums) && nums[tailIndex] == 0 {
// 				tailIndex--
// 			}
// 		}
// 		if prevTotal == tailTotal && tailIndex-prevIndex == 2 {
// 			return prevIndex + 1
// 		}
// 	}
// 	fmt.Println("___________________________")
// 	prevIndex, tailIndex, prevTotal, tailTotal = -1, len(nums), 0, 0
// 	for tailIndex-prevIndex > 2 {
// 		if prevTotal > tailTotal {
// 			prevIndex++
// 			prevTotal += nums[prevIndex]
// 		} else {
// 			tailIndex--
// 			tailTotal += nums[tailIndex]
// 		}
// 		fmt.Println("prevIndex, tailIndex, prevTotal, tailTotal", prevIndex, tailIndex, prevTotal, tailTotal)
// 		for tailIndex-prevIndex > 3 &&
// 			((prevIndex > -1 && nums[prevIndex] == 0) || (tailIndex < len(nums) && nums[tailIndex] == 0)) {
// 			fmt.Println("+++++++++++= prevIndex, tailIndex, prevTotal, tailTotal", prevIndex, tailIndex, prevTotal, tailTotal)
// 			if prevIndex > -1 && nums[prevIndex] == 0 {
// 				prevIndex++
// 			} else if tailIndex < len(nums) && nums[tailIndex] == 0 {
// 				tailIndex--
// 			}
// 		}
// 		if prevTotal == tailTotal && tailIndex-prevIndex == 2 {
// 			return prevIndex + 1
// 		}
// 	}
// 	return -1
// }

// @lc code=end
