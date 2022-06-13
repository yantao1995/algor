package leetcode

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
// func threeSum(nums []int) [][]int {
// 	if len(nums) < 3 {
// 		return nil
// 	}
// 	QuickSort(nums, 0, len(nums)-1)
// 	if nums[0] == 0 && nums[len(nums)-1] == 0 {
// 		return [][]int{{0, 0, 0}}
// 	}
// 	rtn := [][]int{}
// 	if nums[0] <= 0 && nums[len(nums)-1] <= 0 || nums[0] >= 0 && nums[len(nums)-1] >= 0 { //所有元素在0的左或者右,0的单边直接返回
// 		return nil
// 	}
// 	//	fmt.Println(nums)
// 	apdFlag := true
// 	prev, prevCursor := 0, 0
// 	tail, tailCursor := len(nums)-1, len(nums)-1
// 	for prev = 0; nums[prev] < 0; prev++ { //左1，右2
// 		temp := 0 - nums[prev]
// 		for tail = len(nums) - 1; nums[tail] > 0; tail-- {
// 			//fmt.Println(tail, nums[tail], nums[tail] <= temp)
// 			if nums[tail] <= temp {
// 				for tailCursor = tail - 1; nums[tailCursor] >= 0; tailCursor-- {
// 					//fmt.Println(temp, nums[tail], nums[tailCursor])
// 					if temp > nums[tailCursor]+nums[tail] {
// 						goto comeHereLeft
// 					}
// 					if temp == nums[tailCursor]+nums[tail] { //去0重   -1，0，0，1
// 						apdFlag = true
// 						for kk := range rtn {
// 							if nums[prev] == rtn[kk][0] && nums[tailCursor] == rtn[kk][1] &&
// 								nums[tail] == rtn[kk][2] {
// 								apdFlag = false
// 								break
// 							}
// 						}
// 						if apdFlag {
// 							rtn = append(rtn, []int{nums[prev], nums[tailCursor], nums[tail]})
// 						}
// 					}
// 				}
// 			}
// 		comeHereLeft:
// 			if nums[prev+1] == nums[prev] { //去重
// 				prev++
// 			}
// 		}
// 	}
// 	if len(nums) == 3 {
// 		return rtn
// 	}
// 	if prev+2 < len(nums) {
// 		if nums[prev] == 0 && nums[prev+1] == 0 && nums[prev+2] == 0 {
// 			rtn = append(rtn, []int{0, 0, 0})
// 		}
// 	}
// 	for tail = len(nums) - 1; nums[tail] > 0; tail-- { //左2，右1
// 		temp := 0 - nums[tail]
// 		for prev = 0; nums[prev] < 0; prev++ {
// 			if nums[prev] >= temp {
// 				for prevCursor = prev + 1; nums[prevCursor] <= 0; prevCursor++ {
// 					if temp < nums[prev]+nums[prevCursor] {
// 						goto comeHereRight
// 					}
// 					if temp == nums[prev]+nums[prevCursor] {
// 						apdFlag = true
// 						for kk := range rtn {
// 							if nums[prev] == rtn[kk][0] && nums[prevCursor] == rtn[kk][1] &&
// 								nums[tail] == rtn[kk][2] {
// 								apdFlag = false
// 								break
// 							}
// 						}
// 						if apdFlag {
// 							rtn = append(rtn, []int{nums[prev], nums[prevCursor], nums[tail]})
// 						}
// 					}
// 				}
// 			}
// 		comeHereRight:
// 			if nums[tail-1] == nums[tail] { //去重
// 				tail--
// 			}
// 		}
// 	}
// 	return rtn
// }
// func QuickSort(nums []int, low, high int) {
// 	if low >= high {
// 		return
// 	}
// 	i := low + 1
// 	j := high
// 	sufftoprev := true
// 	index := low
// 	for i <= j {
// 		if sufftoprev { //坑在前面，哨兵从后往前
// 			if nums[index] > nums[j] {
// 				exchangeIJ(nums, index, j)
// 				index = j
// 				sufftoprev = false
// 			}
// 			j--
// 		} else { //坑在后面，哨兵从前往后
// 			if nums[index] < nums[i] {
// 				exchangeIJ(nums, index, i)
// 				index = i
// 				sufftoprev = true
// 			}
// 			i++
// 		}
// 	}
// 	QuickSort(nums, low, index-1)
// 	QuickSort(nums, index+1, high)
// }
// func exchangeIJ(nums []int, i, j int) {
// 	nums[i], nums[j] = nums[j], nums[i]
// }

// @lc code=end
